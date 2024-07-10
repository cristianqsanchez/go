package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
  cases := []struct {
    Name          string
    Input         interface{}
    ExpectedCalls []string
  }{
    {
      Name: "struct with one string field",
      Input: struct {
        Name string
      }{"Cris"},
      ExpectedCalls: []string{"Cris"},
    },
    {
      Name: "struct with two string fields",
      Input: struct {
        Name string
        City string
      }{"Cris", "Medellin"},
      ExpectedCalls: []string{"Cris", "Medellin"},
    },
    {
      Name: "struct with non string field",
      Input: struct {
        Name string
        Age int
      }{"Cris", 20},
      ExpectedCalls: []string{"Cris"},
    },
    {
      Name: "nested fields",
      Input: Person{
        "Cris",
        Profile{33, "Medellin"},
      },
      ExpectedCalls: []string{"Cris", "Medellin"},
    },
    {
      Name: "pointers to things",
      Input: &Person{
        "Cris",
        Profile{33, "Medellin"},
      },
      ExpectedCalls: []string{"Cris","Medellin"},
    },
    {
      Name: "slices",
      Input: []Profile {
        {33, "Medellin"},
        {34, "New York"},
      },
      ExpectedCalls: []string{"Medellin", "New York"},
    },
  }

  for _, test := range cases {
    t.Run(test.Name, func(t *testing.T) {
      var got []string

      walk(test.Input, func(input string) {
        got = append(got, input)
      })

      if !reflect.DeepEqual(got, test.ExpectedCalls) {
        t.Errorf("got %v, want %v", got, test.ExpectedCalls)
      }
    })
  }

  t.Run("with maps", func(t *testing.T) {
    aMap := map[string]string{
      "Cow": "Moo",
      "Sheep": "Baa",
    }

    var got []string

    walk(aMap, func(input string) {
      got = append(got, input)
    })

    assertContains(t, got, "Moo")
    assertContains(t, got, "Baa")
  })

  t.Run("with channels", func(t *testing.T) {
    aChannel := make(chan Profile)

    go func() {
      aChannel <- Profile{33, "Berlin"}
      aChannel <- Profile{34, "Paris"}
      close(aChannel)
    }()

    var got []string
    want := []string{"Berlin", "Paris"}

    walk(aChannel, func(input string) {
      got = append(got, input)
    })

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v want %v", got, want)
    }
  })

  t.Run("with function", func(t *testing.T) {
    aFunction := func() (Profile, Profile) {
      return Profile{33, "Berlin"}, Profile{34, "Paris"}
    }

    var got []string
    want := []string{"Berlin", "Paris"}

    walk(aFunction, func(input string) {
      got = append(got, input)
    })

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v, want %v", got, want)
    }
  })
}

func assertContains(t testing.TB, haystack []string, needle string) {
  t.Helper()
  contains := false
  for _, x := range haystack {
    if x == needle {
      contains = true
    }
  }

  if !contains {
    t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
  }
}

type Person struct {
  Name string
  Profile Profile
}

type Profile struct {
  Age int
  City string
}

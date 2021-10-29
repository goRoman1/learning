package main

import "testing"

func TestMaxSide(t *testing.T) {
	testTable := []struct{
		envelope Envelope
		expected float64
	} {
			{
				envelope: Envelope{20, 10},
				expected: 20,
			},
			{
				envelope: Envelope{5.1, 5},
				expected: 5.1,
			},
			{
				envelope: Envelope{1, 1.1},
				expected: 1.1,
			},
	}

	for _, testCase := range testTable {
		result := testCase.envelope.maxSide()

		t.Logf("Calling MaxSide(%v), result %f\n",
			testCase.envelope, testTable)

		if result != testCase.expected {
			t.Errorf("Incorrect expected. Expect expected %v but got %v", testCase.expected, result)
		}
	}
}

func TestMinSide(t *testing.T) {
	testTable := []struct{
		envelope Envelope
		expected float64
	} {
		{
			envelope: Envelope{10, 10.1},
			expected: 10,
		},
		{
			envelope: Envelope{0, 0},
			expected: 0,
		},
		{
			envelope: Envelope{11, 10},
			expected: 10,
		},
	}

	for _, testCase := range testTable {
		result := testCase.envelope.minSide()

		t.Logf("Calling MaxSide(%v), result %f\n",
			testCase.envelope, testTable)

		if result != testCase.expected {
			t.Errorf("Incorrect expected. Expect expected %v but got %v", testCase.expected, result)
		}
	}
}

func TestCanBeNested(t *testing.T) {
	testTable := []struct{
		env1, env2 Envelope
		expected   bool
	}{
		{
			env1: Envelope{1, 1},
			env2: Envelope{2, 2},
			expected: true,
		},
		{
			env1:Envelope{10, 5},
			env2: Envelope{6, 4.9},
			expected: false,
		},
		{
			env1: Envelope{9, 5},
			env2: Envelope{8.9, 5.1},
			expected: false,
		},
	}

	for _, testCase := range testTable {
		result := testCase.env1.CanBeNested(testCase.env2)

		t.Logf("Calling (%v)CanBeNested(%v), result %v\n",
			testCase.env1,testCase.env2, testTable)

		if result != testCase.expected {
			t.Errorf("Incorrect expected. Expect expected  %v but got %v -> for input params %v as envelope 1 and %v as envelope 2", testCase.expected, result, testCase.env1, testCase.env2)
		}
	}
}



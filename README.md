# Fizzbuzz REST Server

This is a simple Fizzbuzz REST server implemented in Golang using the Gin framework.

It provides an API endpoint for generating Fizzbuzz sequences and retrieving statistics about the most common requests.

## Features

- Accepts parameters for generating Fizzbuzz sequences.
- Exposes an API endpoint for generating Fizzbuzz sequences.
- Provides statistics about the most common requests.

## Getting Started

1. Clone this repository:

```bash
git clone https://github.com/yourusername/fizzbuzz.git
cd fizzbuzz
```

2. Build the Fizz-Buzz server:

```bash
make run
```

3. Test endpoint fizzbuzz:

```bash
$ curl -s -d '{"int1":3, "int2":5, "limit":15, "str1":"fizz", "str2":"buzz"}' -H "Content-Type: application/json" -X POST http://localhost:8000/fizzbuzz | jq
[
  "1",
  "2",
  "fizz",
  "4",
  "buzz",
  "fizz",
  "7",
  "8",
  "fizz",
  "buzz",
  "11",
  "fizz",
  "13",
  "14",
  "fizzbuzz"
]
```

4. Test stats endpoint

```bash
$ curl -s  http://localhost:8000/stats | jq
{
  "hits": 1,
  "most_common_request": {
    "int1": 3,
    "int2": 5,
    "limit": 15,
    "str1": "fizz",
    "str2": "buzz"
  }
}
```
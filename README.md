# Vowel & Consonant Counter (Go CLI Tool)

A simple Go command-line application that counts vowels and consonants in a given text input.  
Supports both interactive mode and direct CLI input.

---

## Features
- Counts vowels (a, e, i, o, u)
- Counts consonants
- Ignores numbers and special characters
- Interactive mode (continuous input)
- CLI flag mode for quick usage
- Unit tested core logic

---

## Project Structure
```

vowel-consonant-counter/
│
├── main.go
├── go.mod
└── counter/
├── counter.go
└── counter_test.go

````

---

## Installation

```bash
git clone https://github.com/Perfected1/vowel-consonant-counter.git
cd vowel-consonant-counter
go mod tidy
````

---

## Usage

### 1. Interactive Mode

```bash
go run main.go
```

Type input continuously, type `exit` to quit.

---

### 2. CLI Mode (Flag)

```bash
go run main.go -text="I love Go programming"
```

---

## Run Tests

```bash
go test ./...
```

---

## Example Output

```
=== Vowel & Consonant Counter ===

Enter text: Hello World

Results:
Vowels: 3
Consonants: 7
```

---

## Author

Chike Jerry Nnamadim
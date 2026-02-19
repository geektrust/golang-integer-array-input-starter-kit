Golang Interview – String Input Starter Kit

This repository provides a generic starter kit for solving interview-style problems in Go (Golang), where inputs are passed as raw strings via the command line.

It is designed to closely simulate real interview environments where:
	•	Inputs are non-interactive
	•	Parsing is part of the problem
	•	Only correct logic and exact output matter

⸻

📌 Problem Description

Each problem will:
	•	Provide one or more input strings as arguments
	•	Expect you to parse the string(s)
	•	Apply the required business or algorithmic logic
	•	Print the final result to standard output

The exact problem statement, constraints, and output format will be provided separately.

⸻

🔢 Input / Output Format (Generic)

Input

<input string 1>
<input string 2>
...

	•	Inputs are passed as command-line arguments
	•	Each argument is treated as one complete input string

Output

<problem-specific output>

	•	Output must exactly match the expected format
	•	No additional logs or debug statements

⸻

📂 Project Structure

.
├── run.sh               # Builds and runs the Go program with input strings
├── main.go              # Your implementation
└── README.md            # Documentation


⸻

▶️ Running the Solution

Use run.sh to build and execute the program with input strings.

Syntax

./run.sh '<input 1>' '<input 2>' ...

Example

./run.sh '3 Paris one-way' '2 London'

Each quoted value is passed as a single string argument to the Go program.

🛠️ Where to Implement

Write your solution in:

main.go

The main function is the entry point where:
	1.	Input arguments are read from os.Args
	2.	Parsing and validation logic is applied
	3.	Final output is printed using fmt.Print / fmt.Println

⸻

🧾 Output Rules
	•	Print only the final result to stdout
	•	Follow the output format defined in the problem statement
	•	Avoid extra whitespace, logs, or debug prints

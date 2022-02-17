# PA1456-Individuellt

## Compile instructions

*(Requires GO installed on local machine)*

<img src="https://go.dev/blog/gopher/header.jpg" alt="go gopher" width="150" />


<details>
<summary> MacOS </summary>

to compile for MacOS run:

```bash
GOOS=darwin GOARCH=amd64 go build -o main.exe main.go
```
You might need to do chmod +x main.exe before.

</details>

<details>
<summary>Linux</summary>

To compile for Linux run the bash command:

```bash
GOOS=linux GOARCH=amd64 go build -o main.exe main.go
```
</details>


<details>
<summary>Windows</summary>

To compile for windows machines run the bash command: 

```bash
GOOS=windows GOARCH=amd64 go build -o main.exe main.go
```

</details>

---

## Run instructions

## Todo

- Write detailed run instructions
- More information given to user at run time
- Checks on all linter rules whether they already have been done.
- Checks on file endings
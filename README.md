# ProgressBar



## About
Saturday night fun project, utility to add a progress bar to your CLI tool.

## Preview
![Screencast from 2023-09-03 00-39-02](https://github.com/Vishvajeet590/ProgressBar/assets/42716731/be79fb8c-bee7-4863-85b3-f2a739cd5318)

## Usage

```go
func main() {
	bar := cmd.NewProgressBar(10, 200)
	fmt.Printf("Downlading file....")
	for i := 1.0; i <= 10; i++ {
		time.Sleep(time.Second * 1)
		bar.Update(i)
	}
	time.Sleep(time.Second)
	bar.CleanUp()
}
```

## ToDo
- [ ] Add support to have multiple bars
- [ ] Thread safe

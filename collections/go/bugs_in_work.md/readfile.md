

``` golang

	lines := make([]string, 0)
	file, err := os.Open(filename)
	if err != nil {
		return []string{""}, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		// read one line from the file:
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		lines = append(lines, line[:len(line)-1])
	}
	return lines, err
```


reader.ReadString('\n') would also read the \n character.

needs to trim it.
package main

import (
	"os"
)

func main() {
	// f, err := os.Open("example.txt")

	// if err != nil {
	// log the error
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil {
	// log the error
	// 	panic(err)
	// }

	// fmt.Println("fileInfo:", fileInfo)
	// fmt.Println("Name:", fileInfo.Name())
	// fmt.Println("Size:", fileInfo.Size())
	// fmt.Println("Permissions:", fileInfo.Mode())
	// fmt.Println("Last Modified:", fileInfo.ModTime())
	// fmt.Println("Is Directory?:", fileInfo.IsDir())

	// Read File
	// defer f.Close()

	// buf := make([]byte, 13)

	// data, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("data:", data)
	// fmt.Println("buf:", buf)

	// for i := 0; i < len(buf); i++ {
	// 	println("data:", data, string(buf[i]))
	// }

	// if small file then this way better
	// fileData, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(fileData))

	// if file is large

	// Read Folders
	// dir, err := os.Open(".") // .(dot) means current folder
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfos, err := dir.ReadDir(-1)

	// for _, fi := range fileInfos {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	// Create file
	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("bye buddy,  nice to meet you")

	// bytes := []byte("Hello, I am Mahin Hasan.")

	// f.Write(bytes)

	// Read and Write to another File (Streaming Fashion)

	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }

	// writer.Flush()
	// fmt.Println("written to new file successfully!")

	// Delete File

	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}

}

package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		CommandService()
	}()
	go func() {
		defer wg.Done()

		DataService()
	}()

	wg.Wait()
}

func CommandService() {

	conn, err := net.Dial("tcp", "172.29.40.115:8080")
	if err != nil {
		panic(err)
	}

	// data := map[string]any{
	// 	"sCommand":  "CONNECT",
	// 	"sMainCode": "0",
	// 	"sSubCode":  "000",
	// 	"sIp":       "127.0.0.1",
	// 	"sOs":       "linux",
	// }

	// plainjson, err := json.Marshal(data)
	// if err != nil {
	// 	panic(err)
	// }

	// key := make([]byte, 256/8)
	// copy(key, []byte("iinpro"))

	// iv := make([]byte, 256/8/2)

	// bl, err := aes.NewCipher(key)
	// if err != nil {
	// 	panic(err)
	// }

	// padded := PKCS5Padding([]byte(plainjson), bl.BlockSize(), len(plainjson))

	// dest := make([]byte, len(padded))

	// cipher.NewCBCEncrypter(bl, iv).
	// 	CryptBlocks(dest, padded)

	// b64 := base64.StdEncoding.EncodeToString(dest)

	b64 := "zNTgox3QfSbjl9sVS/qpqNA4yZ9tbg3wI+0tzUIpJx2nOCv6N5+QhSG7/jDtbKCujl9QMyiwZ6nEhnye+oB3SW17aTiLU3/BN3WihjzC1VtuwVuNluWnAuEE43BQJC62"

	n, err := conn.Write([]byte(b64))
	if err != nil {
		panic(err)
	}

	slog.Info("write to server", "length", n, "data", string(b64))

	var b = make([]byte, 1)

	for {
		_, err := conn.Read(b)
		if errors.Is(err, io.EOF) {
			return
		}

		if err != nil {
			panic(err)
		}

		fmt.Fprint(os.Stderr, string(b))

		// slog.Info("read from server", "length", n, "data", string(b))
	}
}

func DataService() {

	conn, err := net.Dial("tcp", "dbos.co.kr:10092")
	if err != nil {
		panic(err)
	}

	var b = make([]byte, 1<<10)

	for {
		n, err := conn.Read(b)
		if errors.Is(err, io.EOF) {
			return
		}

		if err != nil {
			panic(err)
		}

		slog.Info("read from server", "length", n, "data", string(b))
	}
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

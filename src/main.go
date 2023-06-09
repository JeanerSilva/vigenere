package main

import (
	"fmt"
	"strings"
)

func main() {

	original := "Quem se interessa por aprender a falar Português já pode contar com um ensino eficiente. Com os nossos métodos conseguimos ensinar, sobretudo alunos iniciantes, por meio de textos práticos, que favorecem a boa leitura e consequente compreensão do que é ensinado."
	//original := "Qual é a velocidade dos seus downloads? Em poucos segundos, o teste do FAST.com faz uma estimativa da velocidade do seu provedor."
	fmt.Printf("Original: %s\n\n", original)

	originalAdaptada := replaceAscii(original)
	fmt.Printf("Original adaptada: %s\n\n", originalAdaptada)

	key := "Teste"

	textoCifrado, err := cipher(originalAdaptada, key)

	if err == nil {
		fmt.Println("Texto cifrado:", textoCifrado)
	} else {
		fmt.Println(err)
	}

	textoDecifrado, err := decipher(textoCifrado, key)

	if err == nil {
		fmt.Println("\nTexto decifrado:", textoDecifrado)
	} else {
		fmt.Println(err)
	}
}

func cipher(plainTextString string, key string) (string, error) {
	plainText := []byte(plainTextString)
	cipherText := make([]byte, len(plainText))
	keyAlpha := make([]byte, len(key))

	for i := range key {
		if (key[i] >= 'a') && (key[i]) <= 'z' {
			keyAlpha[i] = key[i] - 'a'
		} else if (key[i] >= 'A') && (key[i] <= 'Z') {
			keyAlpha[i] = key[i] - 'A'
		} else {
			return "", fmt.Errorf("Key '%s' is not valid. Use [a-z] or [A-Z]", key)
		}
	}

	idxKey := 0
	for idx := range plainText {
		if lowerCaseKey(plainText[idx]) {
			cipherText[idx] = ((plainText[idx]-'a')+keyAlpha[idxKey])%26 + 'a'
			idxKey = (idxKey + 1) % len(keyAlpha)
		} else if upperCaseKey(plainText[idx]) {
			cipherText[idx] = ((plainText[idx]-'A')+keyAlpha[idxKey])%26 + 'A'
			idxKey = (idxKey + 1) % len(keyAlpha)
		} else {
			cipherText[idx] = plainText[idx]
		}
	}

	return string(cipherText), nil
}

func decipher(cipherTextString string, key string) (string, error) {
	cipherText := []byte(cipherTextString)
	decipherText := make([]byte, len(cipherText))
	keyAlpha := make([]byte, len(key))

	for i := range key {
		if (key[i] >= 'a') && (key[i]) <= 'z' {
			keyAlpha[i] = key[i] - 'a'
		} else if (key[i] >= 'A') && (key[i] <= 'Z') {
			keyAlpha[i] = key[i] - 'A'
		} else {
			return "", fmt.Errorf("Key '%s' is not valid. Use [a-z] or [A-Z]", key)
		}
	}

	idxKey := 0
	for idx := range cipherText {
		if lowerCaseKey(cipherText[idx]) {
			decipherText[idx] = ((cipherText[idx]-'a')+26-keyAlpha[idxKey])%26 + 'a'
			idxKey = (idxKey + 1) % len(keyAlpha)
		} else if upperCaseKey(cipherText[idx]) {
			decipherText[idx] = ((cipherText[idx]-'A')+26-keyAlpha[idxKey])%26 + 'A'
			idxKey = (idxKey + 1) % len(keyAlpha)
		} else {
			decipherText[idx] = cipherText[idx]
		}
	}

	return string(decipherText), nil
}

func replaceAscii(result string) string {
	m := map[string]string{
		"ã": "a", "â": "a", "á": "a", "à": "a",
		"ê": "e", "é": "e", "è": "e",
		"í": "i", "ì": "i",
		"õ": "o", "ô": "o", "ó": "o",
		"ú": "u",
		"ç": "c"}

	for v, k := range m {
		result = strings.ReplaceAll(result, v, k)
	}
	return result
}

func lowerCaseKey(idx byte) bool {
	return idx >= 'a' && idx <= 'z'
}

func upperCaseKey(idx byte) bool {
	return idx >= 'A' && idx <= 'Z'
}

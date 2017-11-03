package resources

import (
	"lengine/logger"
	"os"
	"strconv"
)

type BinaryData struct {
	Bytes  *[]byte
	Length int64
}

func (b *BinaryData) String() string {
	return string((*b.Bytes)[:b.Length])
}

func LoadFile(fileName string) (*os.File, error) {
	f, err := os.Open(fileName)
	if err != nil {
		errorString := "Cannot load file: '" + fileName + "'"
		errorString += "\r\n    Reason: '" + err.Error() + "'"
		logger.Error(errorString)
	}
	return f, err
}

func LoadBytesFromFile(file *os.File) BinaryData {
	stat, err := file.Stat()
	if err != nil {
		errorString := "Cannot get Stat of file: '" + file.Name() + "'"
		errorString += "\r\n    Reason: '" + err.Error() + "'"
		logger.Error(errorString)
	}

	fileSize := stat.Size()
	buffer := make([]byte, fileSize)
	bytesRead, err2 := file.Read(buffer)
	bytesRead64 := int64(bytesRead)
	if err2 != nil {
		errorString := "Cannot Read byte array from file: '" + file.Name() + "'"
		errorString += "\r\n    Reason: '" + err2.Error() + "'"
		logger.Error(errorString)
	}
	if bytesRead64 != fileSize {
		errorString := "Bytes read does not equal fileSize in file: '" + file.Name() + "'"
		errorString += "\r\n    Read:     " + strconv.FormatInt(bytesRead64, 10)
		errorString += "\r\n    Expected: " + strconv.FormatInt(fileSize, 10)
		logger.Error(errorString)
	}

	return BinaryData{Bytes: &buffer, Length: bytesRead64}
}

# Утилита для генерации cookie_hash и cookie_encryption для ручной конфигурации Semaphore

[Конфигурационный файл Semaphore][semaphore_config]

## Использование
```
[WARN] Flags are mutually exclusive, specify at least one of -hash (64 bit) or -encryption (32 bit) [WARN]

Usage of C:\Users\WEBVAL~1\AppData\Local\Temp\go-build3334952565\b001\exe\main.exe:
  -encryption
        generate cookie_encryption
  -hash
        generate cookie_hash
  -help
        show help
```

## Сборка
### Linux
```
go build -o scSemaphore main.go 
chmod +x scSemaphore
```

### Windows
```
go build -o scSemaphore.exe main.go 
```

[semaphore_config]: https://docs.semui.co/administration-guide/configuration#configuration-file

### Пример использования
```
./scSemaphore -hash
+2RxNE4bWG7x8yB293X2kvURYbL22cHOrrJ6OvL0thVbgl6dFA4qOc3BOPuXH/50hr7ix54ap/UdAt/nvPRokg==

./scSemaphore -encryption
VYQDyDm77QWWMnDQfPxN1gbdwMrw60ysbihZ3cJ1rYY=

./scSemaphore -help      

[WARN] Flags are mutually exclusive, specify at least one of -hash (64 bit) or -encryption (32 bit) [WARN]

Usage of C:\Users\webvalera96\Desktop\secure_cookie\scSemaphore.exe:
  -encryption
        generate cookie_encryption
  -hash
        generate cookie_hash
  -help
        show help
```

####Компиляция контрактов
Генерация abi-файлов из исходника контракта

`solc -o abi --abi contract.sol`

Генерация Go пакета из abi-файла

 `abigen --abi=./abi/Contract.abi --pkg=package_name --out=Contract.go`
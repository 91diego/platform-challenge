# Platform Science
This exercise code performs rule-based delivery assignments to drivers through a command that receives 2 files as input, the first with the shipping addresses and the second with the names of the drivers.

First of all you have to run ´go mod tidy´command to download all the necessary dependencies.

## Run test
Execute ´go test -v ./pkg´ commando to execut all the tests inside pkg folder.

## Run command
Execute ´go run main.go readFile -s [file_path] -d [file_path]´ || ´go run main.go readFile --shipmentFile [file_path] --driverFile [file_path]´
You can download the files inside test_files/txt or test_files/csv/ folder to test CLI Command, or if you want you can create your owns.

## Run CLI command in Debbug mode
Create launch.json file in vscode with the next configuration

´{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "readFile",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "main.go",
            "args": [
                "readFile",
                "--shipmentFile=/Users/dgonzalez/Desktop/shipments.txt",
                "--driverFile=/Users/dgonzalez/Desktop/drivers.txt",
            ],
        }
    ]
}´

### Valid args 
--shipmentFile || -s + [file_path] ie: --shipmentFile=[file_path] || -s=[file_path]
--driverFile || -d + [file_path] ie: --driverFile=[file_path] || -d=[file_path]

Then run and debug the project.

### Debbuging VSCode Documentation
- ´https://code.visualstudio.com/docs/editor/debugging´
- ´https://github.com/golang/vscode-go/wiki/debugging´
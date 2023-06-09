# Platform Science
This exercise code performs rule-based delivery assignments to drivers through a command that receives 2 files as input, the first with the shipping addresses and the second with the names of the drivers.

First of all you have to run ´go mod tidy´ command to download all the necessary dependencies for the project.

## Run tests
Execute ´go test -v ./pkg´ commando to run tests.

## Run command
Execute ´go run main.go readFile -s [file_path] -d [file_path]´ || ´go run main.go readFile --shipmentFile [file_path] --driverFile [file_path]´
You can download the files inside test_files/txt or test_files/csv/ folder to test CLI Command, or if you want you can create your owns.

## Run CLI command in Debbug mode
If you prefer, you can run the command through debbuger. For this create launch.json file in vscode with the next configuration:

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

If you need to configure your debugger, you can follow the next links [code.visualstudio](https://code.visualstudio.com/docs/editor/debugging´) or [github.com](https://github.com/golang/vscode-go/wiki/debugging´)

### Valid command args 
--shipmentFile || -s + ´file_path´ ie: --shipmentFile=´file_path´ || -s=[´file_path´
--driverFile || -d + ´file_path´ie: --driverFile=´file_path´|| -d=´file_path´

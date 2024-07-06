Write-Output "Testing sink package..."
go test ./sink/ -v
Write-Output "==================================="
Write-Output "Testing formatter package..."
go test ./formatter/ -v
Write-Output "==================================="

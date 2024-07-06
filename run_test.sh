#!/bin/sh

echo "Testing sink package..."
go test ./sink/ -v
echo "==================================="
echo "Testing formatter package..."
go test ./formatter/ -v
echo "==================================="

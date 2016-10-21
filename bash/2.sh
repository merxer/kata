#!/bin/bash

#Define bash global variable
VAR="global variable"

function bash {
  local VAR="local variable"
  echo $VAR
}

echo $VAR

bash
echo $VAR

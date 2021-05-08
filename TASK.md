# Task

## Thoughts

Man page: https://man7.org/linux/man-pages/man5/crontab.5.html

Cron expression: 
  - valid time elements: 
    - Star   e.g. *
    - Number e.g. 0
    - Range  e.g. 0-9
    - Step   e.g. NUMBER or RANGE or STAR/NUMBER - */15
  - can chain them with commas e.g. 0,5,20-25
  - different elements have diff rules: high, low e.g. min=0,59 hour=0,23 daymth=1-31 mth=1-12 daywk=0-6
  - alternatives valid for daywk, mth e.g. MON-SUN, JAN-DEC - convert to int

Taking me down lexer/parser/AST/FSM thoughts...

## Planning

[x] skaffolding/CLI (Cobra)
[x] lexer
[x] parser
[x] render example
[] validation? errors?
[] push GH and send
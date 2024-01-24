@echo off
set start_time=%time%
for /l %%x in (1, 1, 100) do (
    set NODE_ID=3001
    blockchain.exe send -from 1HqzmYL8UKpw5PRkFjbjkWiYK428tC9mgs -to 18w3VGg4DyEYM5D181nq8VfUGWXXD4QEoW -amount 1
)
set end_time=%time%
echo Start Time: %start_time%
echo End Time: %end_time%
echo Elapsed Time: %time%

pause

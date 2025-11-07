# SEC Pick 'em 2025 Week 11
Week 11 SEC Pick 'em matchups are as follows:

| Team A |  Team B |
|-----------|-----------|
| Georgia | Mississippi State |
| The Citadel | Ole Miss |
| Texas A&M | Missouri |
| Auburn | Vanderbilt |
| Florida | Kentucky |
| LSU | Alabama |


Winners and scores are calculated by finding the xxHash of the teams names and then taking the modulo 74 of the hash to find the score for each team. Mod 74 was selected based on the highest-scoring modern game. Texas A&M beat LSU 74–72 in a seven-overtime game in 2018.

[xxHash Reference](https://xxhash.com/)
# Results
```
go run main.go
Georgia vs Mississippi State: 46 - 56
Mississippi State wins!

The Citadel vs Ole Miss: 2 - 13
Ole Miss wins!

Texas A&M vs Missouri: 29 - 23
Texas A&M wins!

Auburn vs Vanderbilt: 61 - 37
Auburn wins!

Florida vs Kentucky: 32 - 49
Kentucky wins!

LSU vs Alabama: 32 - 62
Alabama wins!

```
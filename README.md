# Stale Data After Database Update Using GORM in Go

This repository demonstrates a common issue encountered when using GORM in Go:  stale data after updating database records.  The `Updates` method in GORM doesn't automatically refresh the object's fields after an update, leading to inconsistencies if the application relies on the updated values directly.

## Bug Description
The `UpdateUser` function showcases this issue. After updating the user in the database, the function doesn't reload the updated data, resulting in the application working with outdated information. 

## Solution
The `bugSolution.go` file demonstrates a robust solution. It explicitly reloads the user from the database after the update, ensuring the application always uses the most current data.
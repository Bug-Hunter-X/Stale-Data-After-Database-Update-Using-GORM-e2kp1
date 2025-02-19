func UpdateUser(user *User) error {
  // ... some code ...
  if err := db.Model(user).Updates(user).Error; err != nil {
    return fmt.Errorf("failed to update user: %w", err)
  }
  // Reload the user from the database
  if err := db.First(user, user.ID).Error; err != nil {
    return fmt.Errorf("failed to reload user: %w", err)
  }
  return nil
}

// ... elsewhere ...

err := UpdateUser(&user)
if err != nil {
  log.Println("Error updating user:", err)
} 
// Now the 'user' variable will reflect the updated data from the database.
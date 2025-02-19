func UpdateUser(user *User) error {
  // ... some code ...
  if err := db.Model(user).Updates(user).Error; err != nil {
    return fmt.Errorf("failed to update user: %w", err)
  }
  return nil
}

// ... elsewhere ...

err := UpdateUser(&user)
if err != nil {
  log.Println("Error updating user:", err)
}
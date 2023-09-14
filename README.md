# gophkeeper

**users**                       
| user_id | login | password | last_update_at | metainfo |
|---------|-------|----------|----------------|----------|
| PRIMARY KEY INT| VARCHAR(256) | VARCHAR(256) | TIMESTAMPTZ | VARCHAR(256)) |

**textdata**                                                
| text_data_id | user_id | text_data | uploaded_at | metainfo |
|--------------|---------|-----------|-------------|----------|
| PRIMARY KEY INT | INT FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE | VARCHAR(256) | TIMESTAMPTZ | VARCHAR(256) |

**binarydata**                                               
| binary_data_id | user_id | binary_data| uploaded_at | metainfo |                                
|----------------|---------|------------|-------------|----------|                  
| PRIMARY KEY INT | INT FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE | VARCHAR(256) | TIMESTAMPTZ | VARCHAR(256) |

**carddata**
| card_data_id | user_id | card_number | card_name | card_date | cvv | uploaded_at | metainfo |
|--------------|---------|-------------|-----------|-----------|-----|-------------|----------|
| PRIMARY KEY INT | INT FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE | VARCHAR(256) | VARCHAR(256) | VARCHAR(256) | VARCHAR(256)) | TIMESTAMPTZ | VARCHAR(256) |

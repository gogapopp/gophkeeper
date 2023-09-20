# gophkeeper

**users**                       
| user_id | login | password | last_update_at | metainfo |
|---------|-------|----------|----------------|----------|
| PRIMARY KEY INT| BYTEA | BYTEA | TIMESTAMPTZ | BYTEA |

**textdata**                                                
| text_data_id | user_id | text_data | uploaded_at | metainfo |
|--------------|---------|-----------|-------------|----------|
| PRIMARY KEY INT | INT FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE | BYTEA | TIMESTAMPTZ | BYTEA |

**binarydata**                                               
| binary_data_id | user_id | binary_data| uploaded_at | metainfo |                                
|----------------|---------|------------|-------------|----------|                  
| PRIMARY KEY INT | INT FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE | BYTEA | TIMESTAMPTZ | BYTEA |

**carddata**
| card_data_id | user_id | card_number | card_name | card_date | cvv | uploaded_at | metainfo |
|--------------|---------|-------------|-----------|-----------|-----|-------------|----------|
| PRIMARY KEY INT | INT FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE | BYTEA | BYTEA | BYTEA | BYTEA | TIMESTAMPTZ | BYTEA |

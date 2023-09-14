# gophkeeper

**users**                       
| user_id | login | password | metainfo |
|----|-------|----------|----------|
| PRIMARY KEY INT| VARCHAR(256) | VARCHAR(256) | VARCHAR(256)) |

**textdata**                                                
| text_data_id | user_id | text_data | metainfo |
|--------------|---------|-----------|----------|
| PRIMARY KEY INT | INT FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE | VARCHAR(256)  | VARCHAR(256) |

**binarydata**                                               
| binary_data_id | user_id | binary_data| metainfo |                                
|----------------|---------|------------|----------|                     
| PRIMARY KEY INT | INT FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE | VARCHAR(256) | VARCHAR(256) |

**carddata**
| card_data_id | user_id | card_number | card_name | card_date | cvv | metainfo |
|--------------|---------|-------------|-----------|-----------|-----|----------|
| PRIMARY KEY INT | INT FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE | VARCHAR(256) | VARCHAR(256) | VARCHAR(256) | VARCHAR(256)) | VARCHAR(256) |

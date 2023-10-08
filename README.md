# gophkeeper
Gophkeeper: this is a storage application for your data with server storage that can be synchronized with two different clients in the same user "account" (all data in storages is encrypted)  

how to setup:  
download the source code and write the "make" command in the console  

TODO: покрытие кода тестами  

WARNING: buttons can be activated twice in one click (this is a tview bug i think)  
In the CardData form the card number is checked by the Luhn algorithm  
In the BinaryData form you need to enter file path and maximum file size up to 1gb  

![project structure](image.png)

**users**                       
| user_id | login | password | user_phrase | last_update_at | metainfo |
|---------|-------|----------|-------------|----------------|----------|
| PRIMARY KEY INT| BYTEA | BYTEA | varchar(128) | TIMESTAMPTZ | BYTEA |

**textdata**                                                
| text_data_id | user_id | unique_key | text_data | uploaded_at | metainfo |
|--------------|---------|------------|-----------|-------------|----------|
| PRIMARY KEY INT | INT FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE | varchar(128) | BYTEA | TIMESTAMPTZ | BYTEA |

**binarydata**                                               
| binary_data_id | user_id | unique_key | binary_data| uploaded_at | metainfo |                                
|----------------|---------|------------|------------|-------------|----------|                  
| PRIMARY KEY INT | INT FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE | varchar(128) | BYTEA | TIMESTAMPTZ | BYTEA |

**carddata**
| card_data_id | user_id | unique_key | card_number | card_name | card_date | cvv | uploaded_at | metainfo |
|--------------|---------|------------|-------------|-----------|-----------|-----|-------------|----------|
| PRIMARY KEY INT | INT FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE | varchar(128) | BYTEA | BYTEA | BYTEA | BYTEA | TIMESTAMPTZ | BYTEA |
 
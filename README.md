# gophkeeper
Gophkeeper: this is a storage application for your data with server storage that can be synchronized with two different clients in the same user "account" (all data in storages is encrypted)    

how to setup:  
download the source code and write the "make" command in the console and you will get .exe files in the cmd directory  

WARNING: buttons can be activated twice in one click (this is a tview bug i think)  
In the card data form the card number is checked by the Luhn algorithm  
In the binary data form you need to enter file path and maximum file size up to 1gb  

project structure  
[![image.png](https://i.postimg.cc/XNfSdVV9/image.png)](https://postimg.cc/c6LbNGtJ)  
registration in the client
[![client-reg.png](https://i.postimg.cc/1340Fck2/client-reg.png)](https://postimg.cc/XrSySF7k)  
client capabilities  
[![client-abl.png](https://i.postimg.cc/htHmdFbQ/client-abl.png)](https://postimg.cc/G4xtW5fb)  
the forms of data that the client supports  
[![client-save.png](https://i.postimg.cc/50jMdrTp/client-save.png)](https://postimg.cc/G80fvXcs)  
saving to the data  
[![client-save-card.png](https://i.postimg.cc/zv56DPP9/client-save-card.png)](https://postimg.cc/BLYMMN3g)  
getting the data  
[![client-get-text.png](https://i.postimg.cc/wvFcJfBK/client-get-text.png)](https://postimg.cc/dkkyc9dn)  

The server is a regular console application with a logger  

**users**                       
| user_id | login | password | user_phrase |
|---------|-------|----------|-------------|
| PRIMARY KEY INT| BYTEA | BYTEA | varchar(128) |

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
 

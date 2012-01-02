`
TINYINT         //[0 to 255 if UNSIGNED]
SMALLINT        //[0 to 65,535]
MEDIUMINT       //[0 to 16,777,215]
INT             //[0 to 4.294E+9]
BIGINT          //[0 to 18.45E+18]
FLOAT32         //p=0-24  --> "FLOAT"
FLOAT64         //Min=+/-1.175E-38
DOUBLE[(M,D)]   //Min=+/-2.225E-308
DECIMAL[(M,[D])]//range
BIT             //zero or converting with
CHAR            //FIXED.
VARCHAR         //M=0-255 <v5.0.3
TINYTEXT        //0-255 Characters
TEXT            //0-65,535 Char's
MEDIUMTEXT      //0-16,777,215 Char's
LONGTEXT        //0-4,294,967,295 Char's
BINARY          //M=0-255 bytes, FIXED.
VARBINARY       //M=0-255 <v5.0.3
TINYBLOB        //0-255 bytes
BLOB            //0-65,535 bytes
MEDIUMBLOB      //0-16,777,215 bytes
LONGBLOB        //0-4,294,967,295 bytes
ENUM            //Column is exactly 1 of
SET             //Column is 0 or more
DATE            //"9999-12-31"
DATETIME        //"1000-01-01 00:00:00" -
TIME            //"838:59:59"
TIMESTAMP       //19700101000000 -
YEAR            //1900 - 2155
`


const (
	TINYINT         //[0 to 255 if UNSIGNED]
	SMALLINT        //[0 to 65,535]
	MEDIUMINT       //[0 to 16,777,215]
	INT             //[0 to 4.294E+9]
	BIGINT          //[0 to 18.45E+18]
	FLOAT32         //p=0-24  --> "FLOAT"
	FLOAT64         //Min=+/-1.175E-38
	DOUBLE[(M,D)]   //Min=+/-2.225E-308
	DECIMAL[(M,[D])]//range
	BIT             //zero or converting with
	CHAR            //FIXED.
	VARCHAR         //M=0-255 <v5.0.3
	TINYTEXT        //0-255 Characters
	TEXT            //0-65,535 Char's
	MEDIUMTEXT      //0-16,777,215 Char's
	LONGTEXT        //0-4,294,967,295 Char's
	BINARY          //M=0-255 bytes, FIXED.
	VARBINARY       //M=0-255 <v5.0.3
	TINYBLOB        //0-255 bytes
	BLOB            //0-65,535 bytes
	MEDIUMBLOB      //0-16,777,215 bytes
	LONGBLOB        //0-4,294,967,295 bytes
	ENUM            //Column is exactly 1 of
	SET             //Column is 0 or more
	DATE            //"9999-12-31"
	DATETIME        //"1000-01-01 00:00:00" -
	TIME            //"838:59:59"
	TIMESTAMP       //19700101000000 -
	YEAR            //1900 - 2155
)



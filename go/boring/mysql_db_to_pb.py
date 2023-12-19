import re

from sqlalchemy import create_engine, Column as SQLColumn, String
from sqlalchemy.orm import sessionmaker
from sqlalchemy.orm.decl_api import declarative_base

Base = declarative_base()


class Table(Base):
    __tablename__ = 'TABLES'
    TABLE_NAME = SQLColumn('TABLE_NAME', String, primary_key=True)
    TABLE_SCHEMA = SQLColumn('TABLE_SCHEMA', String)


class MyColumn(Base):
    __tablename__ = 'COLUMNS'
    COLUMN_NAME = SQLColumn('COLUMN_NAME', String, primary_key=True)
    DATA_TYPE = SQLColumn('DATA_TYPE', String)
    COLUMN_COMMENT = SQLColumn('COLUMN_COMMENT', String)
    TABLE_NAME = SQLColumn('TABLE_NAME', String)
    TABLE_SCHEMA = SQLColumn('TABLE_SCHEMA', String)


def generate_message_structure(db_session, db_name, table):
    columns = db_session.query(MyColumn).filter(MyColumn.TABLE_SCHEMA == db_name, MyColumn.TABLE_NAME == table).all()

    txt = ""
    for index, col in enumerate(columns):
        if col.COLUMN_COMMENT:
            txt += f"    {pb_type(col.DATA_TYPE)} {camel_case(col.COLUMN_NAME)} = {index + 1}; //{col.COLUMN_COMMENT}\n"
        else:
            txt += f"    {pb_type(col.DATA_TYPE)} {camel_case(col.COLUMN_NAME)} = {index + 1};\n"
    txt = f"message {tb_name(table)} {{\n{txt}}}\n"
    return txt


def tb_name(table_name):
    words = re.split(r'_', table_name)
    return ''.join(word.capitalize() for word in words)


def camel_case(s):
    parts = re.split(r'_', s)
    return parts[0].lower() + ''.join(word.capitalize() for word in parts[1:])


def pb_type(mysql_type):
    if mysql_type in ["varchar", "char", "text", "mediumtext", "longtext"]:
        return "string"
    elif mysql_type in ["timestamp", "datetime", "date", "time"]:
        return "google.protobuf.Timestamp"
    elif mysql_type == "bigint":
        return "int64"
    elif mysql_type in ["int", "mediumint", "smallint", "tinyint"]:
        return "int32"
    elif mysql_type in ["double", "decimal"]:
        return "double"
    elif mysql_type == "float":
        return "float"
    elif mysql_type == "json":
        return "google.protobuf.Any"
    elif mysql_type in ["enum", "set"]:
        return "string"
    elif mysql_type in ["binary", "varbinary", "blob", "longblob", "mediumblob", "tinyblob"]:
        return "bytes"
    else:
        return "string"


def generate_proto_file(db_session, db_name):
    tables = db_session.query(Table).filter(Table.TABLE_SCHEMA == db_name).all()

    message = """syntax = "proto3";
package pbdef;
import "google/protobuf/any.proto";
import "google/protobuf/timestamp.proto";

    """
    for table in tables:
        message += generate_message_structure(db_session, db_name, table.TABLE_NAME) + "\n"

    with open("out_py.proto", "w", encoding='utf-8') as file:
        file.write(message)


if __name__ == "__main__":
    db_name = "dopai"
    db_url = "mysql+mysqlconnector://lzj:123456@192.168.31.54:3306/information_schema"

    engine = create_engine(db_url)
    # Base.metadata.create_all(engine)

    Session = sessionmaker(bind=engine)
    db_session = Session()
    generate_proto_file(db_session, db_name)

    print("Protobuf structure has been saved to out.proto.")

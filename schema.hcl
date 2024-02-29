table "books" {
  schema = schema.main
  column "id" {
    null           = false
    type           = integer
    auto_increment = true
  }
  column "isbn" {
    null = false
    type = text
  }
  column "title" {
    null = false
    type = text
  }
  column "language" {
    null = false
    type = text
  }
  column "author" {
    null = false
    type = text
  }
  column "publisher" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.id]
  }
}
schema "main" {
}

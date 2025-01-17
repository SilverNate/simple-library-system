CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    published_date TIMESTAMP NOT NULL,
    "createdAt" timestamp DEFAULT now() NULL,
    "updatedAt" timestamp DEFAULT now() NULL,
    "deletedAt" timestamp NULL
);

CREATE TABLE borrowers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    "createdAt" timestamp DEFAULT now() NULL,
    "updatedAt" timestamp DEFAULT now() NULL,
    "deletedAt" timestamp NULL
);

-- HistoryBorrower Table with Foreign Keys and Constraints
CREATE TABLE history_borrowers (
    id SERIAL PRIMARY KEY,
    borrower_id BIGINT NOT NULL,
    book_id BIGINT NOT NULL,
    borrow_date TIMESTAMP NOT NULL,
    return_date TIMESTAMP,
    due_date TIMESTAMP NOT NULL,
    status varchar(50) NULL,
    "createdAt" timestamp DEFAULT now() NULL,
    "updatedAt" timestamp DEFAULT now() NULL,
    "deletedAt" timestamp NULL

    CONSTRAINT fk_borrower FOREIGN KEY (borrower_id)
        REFERENCES borrowers(id) ON DELETE CASCADE ON UPDATE CASCADE,

    CONSTRAINT fk_book FOREIGN KEY (book_id)
        REFERENCES books(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Indexes
CREATE INDEX idx_history_borrower_status ON history_borrowers (status);

-- Insert Books
INSERT INTO books (title, author, published_date) VALUES
('Effective Java', 'Joshua Bloch', '2001-09-23 00:00:00.000'),
('Clean Code', 'Robert C. Martin', '2001-02-23 00:00:00.000');

-- Insert Borrowers
INSERT INTO borrowers (name, email) VALUES
('Alice Johnson', 'test@gmail.com'),
('Bob Smith', 'test01@gmail.com');

-- Insert History Borrower Data
INSERT INTO history_borrowers (borrower_id, book_id, borrow_date, due_date , status)
VALUES
(1, 1, '2025-01-01 10:00:00', '2025-01-10 10:00:00', 'on time'),
(2, 2, '2025-01-02 12:00:00', '2025-01-12 12:00:00', 'late');
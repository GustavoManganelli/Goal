CREATE DATABASE goaldb;

BEGIN;

DO $$
BEGIN
    CREATE TYPE "PixKeyType" AS ENUM (
        'CPF',
        'CNPJ',
        'Phone',
        'Email',
        'Random'
    );
EXCEPTION
    WHEN duplicate_object THEN NULL;
END
$$;

DO $$
BEGIN
    CREATE TYPE "GroupMemberRole" AS ENUM (
        'CAN_CREATE_DEBT',
        'CANNOT_CREATE_DEBT'
    );
EXCEPTION
    WHEN duplicate_object THEN NULL;
END
$$;

CREATE TABLE IF NOT EXISTS "AuditableEntity" (
    "CreatedBy" VARCHAR(150) NOT NULL,
    "CreatedAt" TIMESTAMP NOT NULL,
    "UpdatedBy" VARCHAR(150),
    "UpdatedAt" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "Users" (
    "Id" SERIAL PRIMARY KEY,
    "FirstName" VARCHAR(100) NOT NULL,
    "LastName" VARCHAR(100) NOT NULL,
    "BirthDate" DATE,
    "Email" VARCHAR(255) NOT NULL,
    "PasswordHash" TEXT NOT NULL
) INHERITS ("AuditableEntity");

CREATE TABLE IF NOT EXISTS "Groups" (
    "Id" SERIAL PRIMARY KEY,
    "Name" VARCHAR(255)
) INHERITS ("AuditableEntity");

CREATE TABLE IF NOT EXISTS "Banks" (
    "Id" SERIAL PRIMARY KEY,
    "Name" VARCHAR(150) NOT NULL
) INHERITS ("AuditableEntity");

CREATE TABLE IF NOT EXISTS "GroupExpenses" (
    "Id" SERIAL PRIMARY KEY,
    "Subject" VARCHAR(255) NOT NULL,
    "TotalAmount" NUMERIC NOT NULL,
    "GroupId" INTEGER NOT NULL
        CONSTRAINT "FK_GroupExpenses_Group"
            REFERENCES "Groups" ("Id")
            ON DELETE NO ACTION ON UPDATE NO ACTION,
    "UserId" INTEGER NOT NULL
        CONSTRAINT "FK_GroupExpenses_User"
            REFERENCES "Users" ("Id")
            ON DELETE NO ACTION ON UPDATE NO ACTION
) INHERITS ("AuditableEntity");

CREATE TABLE IF NOT EXISTS "UserDebts" (
    "Id" SERIAL PRIMARY KEY,
    "Amount" NUMERIC NOT NULL,
    "UserId" INTEGER NOT NULL
        CONSTRAINT "FK_UserDebts_User"
            REFERENCES "Users" ("Id")
            ON DELETE NO ACTION ON UPDATE NO ACTION,
    "GroupExpenseId" INTEGER NOT NULL
        CONSTRAINT "FK_UserDebts_GroupExpense"
            REFERENCES "GroupExpenses" ("Id")
            ON DELETE NO ACTION ON UPDATE NO ACTION,
    "Status" BOOLEAN
) INHERITS ("AuditableEntity");

CREATE TABLE IF NOT EXISTS "GroupMembers" (
    "Id" SERIAL PRIMARY KEY,
    "UserId" INTEGER NOT NULL
        CONSTRAINT "FK_GroupMembers_User"
            REFERENCES "Users" ("Id")
            ON DELETE NO ACTION ON UPDATE NO ACTION,
    "GroupId" INTEGER NOT NULL
        CONSTRAINT "FK_GroupMembers_Group"
            REFERENCES "Groups" ("Id")
            ON DELETE NO ACTION ON UPDATE NO ACTION,
    "Role" "GroupMemberRole" NOT NULL
) INHERITS ("AuditableEntity");

CREATE TABLE IF NOT EXISTS "Accounts" (
    "Id" SERIAL PRIMARY KEY,
    "UserId" INTEGER NOT NULL
        CONSTRAINT "FK_Accounts_User"
            REFERENCES "Users" ("Id")
            ON DELETE NO ACTION ON UPDATE NO ACTION,
    "BankId" INTEGER NOT NULL
        CONSTRAINT "FK_Accounts_Bank"
            REFERENCES "Banks" ("Id")
            ON DELETE NO ACTION ON UPDATE NO ACTION
) INHERITS ("AuditableEntity");

CREATE TABLE IF NOT EXISTS "PixKeys" (
    "Id" Serial PRIMARY KEY,
    "PixKey" TEXT NOT NULL UNIQUE,
    "AccountId" INTEGER NOT NULL
        CONSTRAINT "FK_PixKeys_Account"
            REFERENCES "Accounts" ("Id")
            ON DELETE NO ACTION ON UPDATE NO ACTION,
    "KeyType" "PixKeyType" NOT NULL
) INHERITS ("AuditableEntity");

CREATE TABLE IF NOT EXISTS "Payments" (
    "Id" SERIAL PRIMARY KEY,
    "Amount" NUMERIC NOT NULL,
    "TransactionDate" TIMESTAMP NOT NULL,
    "UserDebtId" INTEGER NOT NULL
        CONSTRAINT "FK_Payments_UserDebt"
            REFERENCES "UserDebts" ("Id")
            ON DELETE NO ACTION ON UPDATE NO ACTION,
    "Proof" TEXT
) INHERITS ("AuditableEntity");

COMMIT;
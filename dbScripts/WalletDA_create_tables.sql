/***********
WalletDA project

Scripts to create database tables
Diogo Andrade 26/08/2021

************/

USE [WalletDA]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

-- Users: primary key for everything else
CREATE TABLE [dbo].[TblUsers](
	[UserId]	[int] NOT NULL IDENTITY(1,1) PRIMARY KEY,
	[UserName] [nvarchar](100) NOT NULL

) ON [PRIMARY]
GO

-- Balance: every user has a "balance" for each kind of CryptoCurrency
CREATE TABLE [dbo].[TblBalance](
	[BalanceId] [int] NOT NULL IDENTITY(1,1) PRIMARY KEY,
	[UserId]	[int] NOT NULL,
	[Currency] [nvarchar](50) NOT NULL,
	[Amount] [real] NOT NULL,

	CONSTRAINT fk_userId_balance_users FOREIGN KEY (UserId) 
    REFERENCES [dbo].[TblUsers](UserId)
) ON [PRIMARY]
GO

-- Transactions: Transaction History. Can be used  
CREATE TABLE [dbo].[TblTransactions](
	[HistoryId] [int] NOT NULL IDENTITY(1,1) PRIMARY KEY,
	[UserId]	[int] NOT NULL,
	[Currency] [nvarchar](50) NOT NULL,
	[Amount] [real] NOT NULL,
	[TimeStamp] [datetimeoffset] NOT NULL,

	CONSTRAINT fk_userId_transactions_users  FOREIGN KEY (UserId) 
    REFERENCES [dbo].[TblUsers](UserId)
) ON [PRIMARY]
GO




USE [WalletDA]
GO

DECLARE	@returnMessage nvarchar(200)

EXEC	[dbo].[spAddTransaction] 'Holland', 'btc',  0.00243, @returnMessage OUTPUT 
EXEC	[dbo].[spAddTransaction] 'Holland', 'eth',  1.235, @returnMessage OUTPUT 
EXEC	[dbo].[spAddTransaction] 'Holland', 'ada',  23.768, @returnMessage OUTPUT 
EXEC	[dbo].[spAddTransaction] 'Garfield', 'xrp',  30.01, @returnMessage OUTPUT 
EXEC	[dbo].[spAddTransaction] 'Garfield', 'doge',  7865645.09876, @returnMessage OUTPUT 
EXEC	[dbo].[spAddTransaction] 'Garfield', 'btc',  1.24, @returnMessage OUTPUT 
EXEC	[dbo].[spAddTransaction] 'McGuire', 'eth',  4.567, @returnMessage OUTPUT 
EXEC	[dbo].[spAddTransaction] 'McGuire', 'ada',  1.0, @returnMessage OUTPUT 
EXEC	[dbo].[spAddTransaction] 'McGuire', 'xrp',  2.54, @returnMessage OUTPUT 
EXEC	[dbo].[spAddTransaction] 'ben', 'btc',  5.12, @returnMessage OUTPUT 



GO

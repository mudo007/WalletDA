-- ================================================
-- Template generated from Template Explorer using:
-- Create Procedure (New Menu).SQL
--
-- Use the Specify Values for Template Parameters 
-- command (Ctrl-Shift-M) to fill in the parameter 
-- values below.
--
-- This block of comments will not be included in
-- the definition of the procedure.
-- ================================================
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		Diogo Andrade
-- Create date: 26/08/2021
-- Description:	creates a transaction and updates the user balance 
-- =============================================
CREATE PROCEDURE [dbo].[spAddTransaction]
-- Add the parameters for the stored procedure here
@UserName nvarchar(100),
@Currency nvarchar(50),
@Amount real,
@returnMessage nvarchar(200) out

AS
BEGIN

--variable declarations
declare
@userID int,
@hasCurrency bit,
@CurrentAmount real = 0.0,
@newBalance real = 0.0


set @hasCurrency = 0
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;
	-- Wrap the insert on the transaction history and the Balance update into a transaction
	BEGIN TRY
	BEGIN TRANSACTION
		--get the userID
		set @userID =( Select UserId From [dbo].[TblUsers] where UserName = @UserName)
		if @userID is null
		begin
			--crate new user if not found
			INSERT INTO [dbo].[TblUsers] (UserName) VALUES(@UserName);
			--retrieve user id
			set @userID =  SCOPE_IDENTITY()
		end
		--insert transaction
		INSERT INTO [dbo].[TblTransactions] (UserId, Currency, Amount, TimeStamp) Values
			(@userID, @Currency, @Amount, CONVERT(NVARCHAR(30), GETDATE(), 127))

		--check if the user already has a balance for the Currency
		Select 
			@hasCurrency = case when Currency is null then 0  else 1 end ,
			@CurrentAmount = case when Currency is null then 0.0  else Amount end   from TblBalance where UserId = @userID and Currency = @Currency
		
		--What would the new balance be?
		set @newBalance = @CurrentAmount + @Amount
		--refuse Overdraw
		if @newBalance < 0.0
		begin
			RAISERROR('Insufficient funds',16,1);
		end

		--insert or update the balance
		if @hasCurrency = 0
		begin	
			--this is the first transaction for this kind of currency. We should not allow it insert a negative value
			insert  into [dbo].[TblBalance] (UserId, Currency, Amount) VALUES (@userID, @Currency, @Amount)	
		end
		else
		begin
			--must update the already existing balance
			update [dbo].[TblBalance] set Amount = @newBalance where UserId = @userID and Currency = @Currency
		end		
		-- if not error, commit the transcation
        COMMIT TRANSACTION
		set @returnMessage =  'Success'
		END TRY
   BEGIN CATCH
		-- if error, roll back any chanegs done by any of the sql statements
		ROLLBACK TRANSACTION
		set @returnMessage = ERROR_MESSAGE()
  END CATCH

	



END
GO

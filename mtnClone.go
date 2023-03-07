package main
import ("fmt")
//MAIN METHOD TO CALL MY FUNCTIONS
func main() {
	fmt.Print("Enter the MTN code : ")
	var mtnCode string
	fmt.Scanln(&mtnCode)

	switch mtnCode{
	case "*131#" : optionDisplay()
	
	 default :  fmt.Println("Invalid code")
	}
}
//OPTION DISPLAY FUNCTION TO START THE PROGRAM
func optionDisplay(){
	mtnOptions := []string {"1.Data Plans","2.Social Bundles","3.Balance Check","4.Roaming/Int'l","5.Borrow Credit/Recharge","6.Gift Data","7.Video Packs","8.Hot Deals","0.Back"}
	
	for _, val := range mtnOptions {
		fmt.Printf("%v\n", val)
	 }
	 
	 fmt.Print("\nEnter choice : ")
	 var dataOptionsInput int
	 fmt.Scanln(&dataOptionsInput)

	 //SWITCH CASE FOR THE DATA OPTIONS INPUT
	 switch dataOptionsInput {
		//CASE 0 OF DATA OPTIONS INPUT
		case 0 : main()
		//CASE 1 OF DATA OPTIONS INPUT
	 case 1 : dataOptions := []string {"Buy Data Plans","1.Daily","2.Weekly","3.Monthly","4.2-3 Month", "5.Always ON Plans","0.Back"}
	 for _, val := range dataOptions {
		 fmt.Printf("%v\n", val)
	  }
	  fmt.Print("\nEnter Choice : ")
	  var dataPlansOptionInput string
	  fmt.Scanln(&dataPlansOptionInput)

	  //NESTED SWITCH CASE IN CASE 1 OF DATA OPTIONS INPUT
	  switch dataPlansOptionInput {
	  case "0" : optionDisplay()
	  case "1" : fmt.Println("Dear Customer,\nYour Daily data plan 40mb for ₦50 was successful")
	  case "2" : fmt.Println("Dear Customer,\nYour weekly data plan 1gb for ₦200 was successful")
	  case "3" : fmt.Println("Dear Customer,\nYour monthly data plan 4gb for ₦1000 was successful")
	  case "4" : fmt.Println("Dear Customer,\nYour Bi-Month data plan 6gb for ₦1500 was successful")
	  case "5" : fmt.Println("Service temporarily unavailable")
	  default : fmt.Println("Invalid input")
	  
	  }

	  //CASE 2 OF DATA OPTIONS INPUT
	  case 2 : socialBundles := []string {"1.WhatsApp","2.Facebook","3.Instagram","4.TikTok", "5. Ayoba","6.All Social Bundles","7.Youtube, Instagram, and TikTok","8.Opera Mini & News","9.Social Mega Bundle","0.Back"}
	 for _, val := range socialBundles {
		 fmt.Printf("%v\n", val)
	  }
	  fmt.Print("Enter Choice : ")
	  var socialBundlesInput string
	  fmt.Scanln(&socialBundlesInput)
	  //nested switch case inside case 2 of data options input
	  switch socialBundlesInput{
	  case "0" : main()
	  case "1": fmt.Println("Whatsapp Only bundle Activated")
	  case "2": fmt.Println("Facebook only bundle activated")
	  case "3" : fmt.Println("Instagram Only Bundle Activated")
	  case "4" : fmt.Println("Tiktok only bundle activated")
	  case "5" : fmt.Println("Ayoba bundle activated")
	  case "6" : fmt.Println("All social bundles plan activated")
	  case "7" : fmt.Println("Purchase successful\nMobile data is restricted to Youtube, Instagram and Tiktok")
	  case "8" : fmt.Println("News Bundle Activated")
	  case "9" : fmt.Println("Social Mega Bundle plan activated")
	  default : fmt.Println("Invalid Input")
	  }

	  
	  //CASE 3 OF DATA OPTIONS INPUT
	case 3 : balanceCheck := []string{"mPulse Main Account: N650.91 Airtime Bundle: N900.12. Data Balance : 800MB"}
	for _, val := range balanceCheck {
		fmt.Printf("%v\n", val)
	 }
	 //CASE 4 OF DATA OPTIONS INPUT
	case 4: roaming := []string {"Service currently unavailable"}
	for _, val := range roaming {
		fmt.Printf("%v\n", val)
	 }
	 //CASE 5 OF DATA OPTIONS INPUT
	case 5 : borrowCredit := []string{"1.Borrow Mobile data","2.Borrow Credit"}
	for _, val := range borrowCredit{
		fmt.Printf("%v\n", val)
	 }
	 fmt.Println("Enter Choice : ")
	 var borrowInput int
	 fmt.Scanln(&borrowInput)

	 //nested switch case inside case 5 of data options input
	 	switch borrowInput{
			//case 1 of nested switch case
		case 1 : fmt.Print("How many MegaBytes would you like to borrow? : ")
		var mobileDataBorrowInput string
		fmt.Scanln(&mobileDataBorrowInput)
		fmt.Println("You have successfully borrowed "+mobileDataBorrowInput+"MB")
		//case 2 of nested switch case
		case 2 : creditBorrowing := []string{"1.N100","2.200","3.500","4.1000"}
		for _, val := range creditBorrowing {
			fmt.Printf("%v\n", val)
		 }
		fmt.Println("Enter choice : ")
		var creditBorrowInput int
		fmt.Scanln(&creditBorrowInput)
		//nested switch case inside nested switch case inside switch case
		switch creditBorrowInput{
		case 1 : fmt.Println("You have successfully borrowed N100, It would be deducted from your next recharge")
		case 2 : fmt.Println("You have successfully borrowed N200, It would be deducted from your next recharge")
		case 3 : fmt.Println("You have successfully borrowed N500, It would be deducted from your next recharge")
		case 4 : fmt.Println("You have successfully borrowed N1000, It would be deducted from your next recharge")
		default : fmt.Println("Invalid input")
		}

		}
		//CASE 6 OF DATA OPTIONS INPUT
	case 6 : fmt.Println("How many MegaBytes would you like to gift? : ")
	var giftAmount string
	fmt.Scanln(&giftAmount)
	fmt.Println("Enter the number of the receiver : ")
	var giftReceiver string
	fmt.Scanln(&giftReceiver)
	fmt.Println("You have successfully gifted "+giftAmount+"mb to "+giftReceiver)

		//CASE 7 OF DATA OPTIONS INPUT
	case 7 : videoPacks := []string{"Video Streaming Packs","1.Youtube Video Packs","2.StarTimes Video Packs","3.1GB(Youtube Only)+500MB(Data Access)","4.Showmax Mobile"}
	for _, val := range videoPacks {
		fmt.Printf("%v\n", val)
	 }
	 fmt.Print("Enter Choice : ")
	 var videoPacksInput string
	 fmt.Scanln(&videoPacksInput)
	 //nested switch case inside case 7 of data options input
	 switch videoPacksInput {
	 case "1" : fmt.Println("Youtube Pack successfully purchased")
	 case "2" : fmt.Println("StarTimes Pack successfully purchased")
	 case "3" : fmt.Println("Dear Customer\nYou have successfully received 1GB of Data (Youtube Only)")
	 case "4" : fmt.Println("ShowMax Mobile plan successfully purchased")
	 default : fmt.Println("Invalid Input")
	 }
	 //CASE 8 OF DATA OPTIONS INPUT
	case 8 : hotDeals := []string{"1.TopDeal4Me","2.Data4Me","3.Recharge4ME","4.Main Menu"}
	for _, val := range hotDeals {
		fmt.Printf("%v\n", val)
	 }
	 fmt.Print("Enter Choice : ")
	 var hotDealsInput string
	 fmt.Scanln(&hotDealsInput)

	 //nested switch case inside case 8 of data options input
	 switch hotDealsInput{
		//case 1 of hotDealsInput switch case
	 case "1": topDeals := []string{"1.Enjoy 1.5GB for N300","0.Main Menu"}
	 for _, val := range topDeals {
		fmt.Printf("%v\n", val)
	 }
	 fmt.Print("Enter Choice : ")
	 var topDealsInput string
	 fmt.Scanln(&topDealsInput)
	 //nested switch case in case 1 of hotDealsInput switch case
	 switch topDealsInput {
	 case "1" : fmt.Println("You have successfully received 1.5GB for N300\nData is valid for one week")
	 case "0" : optionDisplay()
	 default : fmt.Println("Invalid Input")
	 }
	 
	 //case 2 of hotDealsInput switch case
	case "2" : dataForMe := []string{"1.Enjoy 2GB for N500","2.Enjoy 1GB for N200","0.Main Menu"}
	for _, val := range dataForMe {
		fmt.Printf("%v\n", val)
	 }
	 fmt.Print("Enter Choice : ")
	 var dataForMeInput string
	 fmt.Scanln(&dataForMeInput)
	 //nested switch case in case 2 of hot deals input
	 switch dataForMeInput {
	 case "1" : fmt.Println("You have successfully received 2GB for N500\nData is valid for one week")
	 case "2" : fmt.Println("You have successfully received 1GB for N200\nData is valid for one week")
	 case "0" : optionDisplay()
	 default : fmt.Println("Invalid Input")
	 }
	 //case 3 of hot deals input switch case
	 case "3" : rechargeForMe := []string{"Recharge N1000 and get N6000","Recharge N1000 and get free 800MB","0.Main Menu"}
	for _, val := range rechargeForMe {
		fmt.Printf("%v\n", val)
	 }
	 fmt.Print("Enter Choice : ")
	 var rechargeForMeInput string
	 fmt.Scanln(&rechargeForMeInput)
	 //nested switch case in case 2 of hot deals input
	 switch rechargeForMeInput {
	 case "1" : fmt.Println("Recharge Successful\nYou have received N6000 to call all networks")
	 case "2" : fmt.Println("Recharge Successful\nYou have received 800MB to browse the internet")
	 case "0" : optionDisplay()
	 default : fmt.Println("Invalid Input")
	 }
	 //case 4 of hot deals input
	case "4" : optionDisplay()
	 }
		//DEFAULT CASE OF DATA OPTIONS INPUT
		default : fmt.Println("Invalid Input")
	}
	 }

	
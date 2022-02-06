package getAccountInfo

import "os"

func accountInfo() (string, string) {

	return os.Getenv("ACCOUNT_NAME"), os.Getenv("ACCOUNT_KEY")
	// return "fc29f08dfb238435fbcf4ec", "RayUPDk/VRZ80/2OheDHJWZg76htfobGy5Me2PpsNOKh5m4+Glh0PNNem/YbZ0MRNtUMbxtG7VE6g/hByq9JdQ=="

}

package exercises

import "testing"

func TestBankAccount(t *testing.T) {
	// t.Skip("Chapter 06 exercise: implement BankAccount, then delete this Skip")

	cases := []struct {
		name        string
		start       float64
		op          string // "deposit" or "withdraw"
		amount      float64
		wantBalance float64
		wantErr     bool
	}{
		{"deposit adds", 100, "deposit", 50, 150, false},
		{"deposit negative errors", 100, "deposit", -10, 100, true},
		{"withdraw subtracts", 100, "withdraw", 40, 60, false},
		{"withdraw exact balance", 100, "withdraw", 100, 0, false},
		{"withdraw negative errors", 100, "withdraw", -5, 100, true},
		{"withdraw insufficient errors", 100, "withdraw", 500, 100, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			acc := &BankAccount{Owner: "Test", Balance: tc.start}

			var err error
			switch tc.op {
			case "deposit":
				err = acc.Deposit(tc.amount)
			case "withdraw":
				err = acc.Withdraw(tc.amount)
			}

			if tc.wantErr && err == nil {
				t.Fatalf("%s(%v): got nil error, want an error", tc.op, tc.amount)
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("%s(%v): unexpected error: %v", tc.op, tc.amount, err)
			}
			if acc.Balance != tc.wantBalance {
				t.Errorf("balance = %v, want %v", acc.Balance, tc.wantBalance)
			}
		})
	}
}

func TestBankAccountString(t *testing.T) {
	// t.Skip("Chapter 06 exercise: implement BankAccount.String, then delete this Skip")

	acc := &BankAccount{Owner: "Alice", Balance: 150}
	want := "Alice: 150.00"
	if got := acc.String(); got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}

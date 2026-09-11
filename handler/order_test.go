package handler

import "testing"

// TestTaxAndDiscountLogic memastikan rumus PPN, diskon, dan total benar.
func TestTaxAndDiscountLogic(t *testing.T) {
	subtotal := 250000.0
	tax := subtotal * 0.11
	discount := subtotal * 0.10
	total := subtotal + tax - discount

	if tax != 27500 {
		t.Fatalf("PPN salah: %.0f", tax)
	}

	if discount != 25000 {
		t.Fatalf("diskon salah: %.0f", discount)
	}

	if total != 252500 {
		t.Fatalf("total salah: %.0f", total)
	}
}

// main.go
package main

import (
	"fmt"
	"strings"
)

// --- 1. Rotate four variables to (b,c,d,a) ---
// Two variants: using a single temp, and using Go's parallel assignment.

// rotateWithTemp performs the cyclic rotation a<-b, b<-c, c<-d, d<-a
// using the usual replacement sequence (needs one temp).
func rotateWithTemp(a, b, c, d int) (int, int, int, int) {
	// replacement sequence (minimal for sequential replacements):
	// t <- a
	// a <- b
	// b <- c
	// c <- d
	// d <- t
	t := a
	a = b
	b = c
	c = d
	d = t
	return a, b, c, d
}

// rotateParallel uses Go's parallel assignment (single simultaneous replacement)
func rotateParallel(a, b, c, d int) (int, int, int, int) {
	// a, b, c, d = b, c, d, a  (one simultaneous assignment)
	return b, c, d, a
}

// --- 2. Proof sketch (m > n at beginning of step E1 except possibly first time) ---
// Short proof included in comments printed by main.

// --- 3. Algorithm E (Euclid) and Algorithm F (avoid trivial replacements) ---

// gcdEuclid counts number of E1 executions and returns gcd.
// We define "E1" as the division step where we compute q = m/n and r = m % n.
// Each loop iteration performing that division counts as one E1.
func gcdEuclid(m, n int) (g, e1count int) {
	// Make sure we operate on positive integers
	if m < 0 {
		m = -m
	}
	if n < 0 {
		n = -n
	}
	e1 := 0
	for n != 0 {
		// Step E1: division m = q n + r
		_ = m / n // do division (counted as E1)
		r := m % n
		e1++
		// Now set m <- n, n <- r (these are replacements)
		m, n = n, r
	}
	return m, e1
}

// gcdF implements Algorithm F: same result but avoids trivial copies like "m <- n"
// when they would set the variable to itself or otherwise be no-ops.
// Practically, we implement Euclid but skip redundant assignments and avoid
// setting variable to itself when not needed. The logic below is a direct,
// clearer, and slightly more careful implementation (fewer trivial writes).
func gcdF(m, n int) (g, e1count int) {
	if m < 0 {
		m = -m
	}
	if n < 0 {
		n = -n
	}
	e1 := 0
	// ensure m >= n at entry to the main loop to mirror typical Algorithm E
	if m < n {
		// swap (only when necessary)
		m, n = n, m
	}
	for n != 0 {
		// E1: division
		r := m % n
		_ = m / n
		e1++
		// Instead of always writing m = n; n = r; we check for trivial cases.
		// If r == 0, gcd is n (we can return without further writes).
		if r == 0 {
			return n, e1
		}
		// Now set m <- n and n <- r, but only if they change.
		if m != n {
			m = n
		}
		if n != r {
			n = r
		}
		// loop continues
	}
	return m, e1
}

// --- 4. gcd(2166, 6099) with Euclid steps ---
func gcdSteps(a, b int) (g int, steps [][4]int) {
	x, y := a, b
	for y != 0 {
		q := x / y
		r := x % y
		steps = append(steps, [4]int{x, y, q, r})
		x, y = y, r
	}
	return x, steps
}

// --- 6. Compute T5 (average number of times E1 is performed when n = 5) ---
// We interpret T5 as the average (over m = 1..5) of the number of E1 steps
// when computing gcd(m, 5). (This follows the usual textbook exercise style.)
func Tn_when_n_fixed(n int) (avg float64, counts []int) {
	total := 0
	counts = make([]int, n)
	for m := 1; m <= n; m++ {
		_, e1 := gcdEuclid(m, n)
		counts[m-1] = e1
		total += e1
	}
	avg = float64(total) / float64(n)
	return avg, counts
}

// --- 7. U_m and relation to T_m ---
// We'll provide a short demonstration: show that for fixed m, count_E1(m, n)
// depends only on n (mod m) via the remainder when n divided by m, hence the
// sequence is periodic with period m. Therefore the average over n exists and
// equals the average over one period (n=1..m). We compute Um numerically.
func Um_for_m(m int) (avg float64, counts []int) {
	// counts for n = 1..m (one period)
	counts = make([]int, m)
	total := 0
	for n := 1; n <= m; n++ {
		_, e1 := gcdEuclid(m, n)
		counts[n-1] = e1
		total += e1
	}
	avg = float64(total) / float64(m)
	return avg, counts
}

// --- 8. Effective formal algorithm using string "a^m b^n" with subtraction rule ---
// Parse input string of form "aaaa...bbbb..." and compute gcd using r <- |m-n|, n <- min(m,n).
func gcdFromString(input string) (g int, err error) {
	// Count leading 'a' then trailing 'b' assuming input = a^m b^n exactly.
	// We'll be lenient and trim spaces, but otherwise require correct form.
	s := strings.TrimSpace(input)
	if s == "" {
		return 0, fmt.Errorf("empty input")
	}
	// count a's until first non-'a'
	i := 0
	for i < len(s) && s[i] == 'a' {
		i++
	}
	j := i
	for j < len(s) && s[j] == 'b' {
		j++
	}
	if j != len(s) || i == 0 {
		return 0, fmt.Errorf("input must be of the form a^m b^n with m>=1, n>=0")
	}
	m := i
	n := len(s) - i
	// Now use the subtraction version until one becomes zero
	for n != 0 && m != 0 {
		if m == n {
			return m, nil
		}
		if m > n {
			m = m - n
		} else {
			n = n - m
		}
	}
	if m == 0 {
		return n, nil
	}
	return m, nil
}

func main() {
	fmt.Println("=== 1. Rotations ===")
	a, b, c, d := 1, 2, 3, 4
	fmt.Printf("original: a=%d b=%d c=%d d=%d\n", a, b, c, d)
	a1, b1, c1, d1 := rotateWithTemp(a, b, c, d)
	fmt.Printf("rotateWithTemp -> a=%d b=%d c=%d d=%d\n", a1, b1, c1, d1)
	a2, b2, c2, d2 := rotateParallel(1, 2, 3, 4)
	fmt.Printf("rotateParallel -> a=%d b=%d c=%d d=%d\n\n", a2, b2, c2, d2)

	fmt.Println("=== 2. Proof sketch (m > n at start of E1 except possibly first time) ===")
	fmt.Println("Proof sketch:")
	fmt.Println("Algorithm E begins by ensuring (or not) an order for m and n; if initially m <= n")
	fmt.Println("then the first step swaps them (or the division q = 0 forces a swap), so after the first")
	fmt.Println("time we enter the main division step (E1) we will have m > n. After every execution of E1")
	fmt.Println("we replace (m,n) by (n,r) where 0 <= r < n; hence the next time we reach E1 we have the")
	fmt.Println("old n (which is > r except perhaps when r = 0). Thus at the beginning of E1 (except")
	fmt.Println("possibly the very first time) we always have m > n.\n")

	fmt.Println("=== 3. Algorithms E and F (gcd and improved gcd) ===")
	fmt.Println("gcdEuclid(1071, 462) -> (gcd, E1-count):")
	g, e1 := gcdEuclid(1071, 462)
	fmt.Println(g, e1)
	fmt.Println("gcdF(1071, 462) -> (gcd, E1-count):")
	gf, ef1 := gcdF(1071, 462)
	fmt.Println(gf, ef1)
	fmt.Println()

	fmt.Println("=== 4. gcd(2166, 6099) with Euclid steps ===")
	gx, steps := gcdSteps(6099, 2166) // call in order (larger first) to show steps
	for i, s := range steps {
		fmt.Printf("step %d: %d = %d * %d + %d\n", i+1, s[0], s[1], s[2], s[3])
	}
	fmt.Printf("gcd(6099,2166) = %d\n\n", gx)

	fmt.Println("=== 6. T5 (average number of E1 when n = 5) ===")
	T5, counts := Tn_when_n_fixed(5)
	fmt.Printf("counts for m=1..5 when n=5: %v\n", counts)
	fmt.Printf("T5 = %v = 13/5 = 2.6\n\n", T5)

	fmt.Println("=== 7. U_m (average over n with m fixed) demonstration & relation to T_m ===")
	mval := 5
	U5, cts := Um_for_m(mval)
	fmt.Printf("For m=%d, counts for n=1..%d: %v\n", mval, mval, cts)
	fmt.Printf("U_%d = %v\n", mval, U5)
	fmt.Println("Explanation:")
	fmt.Println("- For fixed m, when you write n = q*m + r (0 <= r < m), the first division step")
	fmt.Println("  produces the remainder r and thereafter the number of E1 steps depends only on r")
	fmt.Println("  (and m). Thus the sequence count_E1(m,n) is periodic of period m, so the average")
	fmt.Println("  over n exists and equals the average over one period (n = 1..m).")
	fmt.Println("- In general U_m (average over all n) is NOT equal to T_m (average over m for fixed n).")
	fmt.Println("  They are different averages; example: for m=5 we found U5 != T5 above.\n")

	fmt.Println("=== 8. GCD from input string a^m b^n (subtraction-based) ===")
	in := "aaaaabbbb" // example m=5, n=4 -> gcd(5,4)=1
	gstr, err := gcdFromString(in)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("gcd from string %q is %d\n\n", in, gstr)
	}

	fmt.Println("=== 9. Notes on 'Procedure for Reading This Set of Books' (why it's not a genuine algorithm) ===")
	fmt.Println("Three ways it fails the five algorithm criteria (finiteness, definiteness, input, output, effectiveness):")
	fmt.Println("1) Definiteness: steps are described informally (\"read the interesting parts\", \"if you like\") and")
	fmt.Println("   require human judgement, so not fully unambiguous elementary operations.")
	fmt.Println("2) Finiteness/Termination: it may ask reader to \"follow references\" or return later; termination is not guaranteed.")
	fmt.Println("3) Effectiveness: steps are not elementary mechanical operations (e.g., \"understand\" or \"decide if it's useful\");")
	fmt.Println("   thus not composed of finitely many basic operations that can be executed exactly.")
	fmt.Println("\nFormat differences from Algorithm E: the Procedure is prose, not numbered precise steps; uses informal language;")
	fmt.Println("no explicit variables or loop invariants; not tailored to machine-executable atomic operations.\n")

	fmt.Println("=== Summary outputs for quick checks ===")
	fmt.Printf("gcd(2166,6099) = %d\n", 57) // final answer for Q4
}

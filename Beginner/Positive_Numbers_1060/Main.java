package Beginner.Positive_Numbers_1060;

import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int amountPositiveNumbers = 0;
        double actualNumber;
        for (int count = 0; count < 6; count ++) {
            actualNumber = sc.nextDouble();
            if (actualNumber >= 0) {
                amountPositiveNumbers += 1;
            }
        }
        System.out.printf("%d valores positivos\n", amountPositiveNumbers);
        sc.close();
    }
    
}

package Beginner.Even_Square_1073;

import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int number = sc.nextInt();
        for (int i = 2; i <= number; i += 2) {
            System.out.printf("%d^2 = %d\n", i, i * i);
        }

        sc.close();
    }
}
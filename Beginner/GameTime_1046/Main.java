import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String[] inputtedValues = sc.nextLine().split(" ");
        int startTime = Integer.parseInt(inputtedValues[0]);
        int endTime = Integer.parseInt(inputtedValues[1]);
        int duration = 24;

        if (endTime > startTime) {
            duration = endTime - startTime;
        }
        else if (endTime < startTime) {
            duration = 24 + endTime - startTime;
        }

        System.out.printf("O JOGO DUROU %d HORA(S)\n", duration);

        sc.close();
    }
}

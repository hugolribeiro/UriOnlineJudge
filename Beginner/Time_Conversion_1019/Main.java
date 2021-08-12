import java.util.Scanner;

public class Main{
    public static void main(String[] args){
        Scanner sc = new Scanner(System.in);
        int durationSeconds = sc.nextInt();
        int hours = durationSeconds / 3600;
        durationSeconds = durationSeconds % 3600;
        int minutes = durationSeconds / 60;
        durationSeconds = durationSeconds % 60;

        System.out.printf("%d:%d:%d\n", hours, minutes, durationSeconds);
        sc.close();
    }
}
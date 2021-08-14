import java.util.Scanner;

public class Main{
    public static void main(String[] args){
        Scanner sc = new Scanner(System.in);
        int personAgeInDays = sc.nextInt();
        int years = personAgeInDays / 365;
        personAgeInDays = personAgeInDays % 365;
        int months = personAgeInDays / 30;
        personAgeInDays = personAgeInDays % 30;

        System.out.println(years + " ano(s)");
        System.out.println(months + " mes(es)");
        System.out.println(personAgeInDays + " dia(s)");
        sc.close();
    }
}
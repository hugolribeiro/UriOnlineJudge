import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Scanner;

public class Main{

    public static void main(String[] args){
        Scanner sc = new Scanner(System.in);
        String[] numbers = sc.nextLine().split(" ");
        List<Integer> numbersInt = new ArrayList<>();
        
        for (String numberStr : numbers) {
            numbersInt.add(Integer.parseInt(numberStr));
        }
        Integer highestNumber = Collections.max(numbersInt);

        System.out.printf("%d eh o maior\n", highestNumber);
    }
}
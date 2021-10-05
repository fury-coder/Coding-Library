import java.util.*;
class insertion
{
    public static void main(String args[])
    {
        Scanner d=new Scanner(System.in);
        int i,j,t,k,min,p;
        int a[]=new int[10];
        for(i=0;i<10;i++)
            a[i]=d.nextInt();
            
        for(i=1;i<10;i++)//1 ->n
        {
            t=a[i];
            j=i-1;
            while(j>=0 && t<a[j])
            {
                a[j+1]=a[j];
                j--;
            }
            a[j+1]=t;
            for(k=0;k<10;k++)
            System.out.print(a[k]+" ");
            System.out.println();
        }
        for(i=0;i<10;i++)
            System.out.print(a[i]+" ");
    }
}

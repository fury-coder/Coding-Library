import java.util.*;
public class mergesort
{
    public static void main(String args[])
    {
        Scanner d=new Scanner(System.in);
        int i;
        int a[]=new int[10];
        for(i=0;i<10;i++)
        a[i]=d.nextInt();
        sort(a,0,10-1);//function calling
    }
    public static void merge(int a[],int l,int m,int r)
    {
        int n1,n2;
        n1=m-l+1;
        n2=r-m;
        int left[]=new int[n1];
        int right[]=new int[n2];
        int i,j;
        for(i=0;i<n1;i++)
        {
            left[i]=a[i];
        }
        for(j=0;j<n2;j++)
        {
            right[j]=a[m+j+1];
        }
        i=j=0;
        int k=0;
        while(i<n1 && j<n2)
        {
           if(left[i]<=right[j])
           {
               a[k]=left[i];
               i++;
            }
            else
            {
                a[k]=right[j];
                j++;
            }
            k++;
        }
        while(i<n1)
        {
            a[k]=left[i];
            i++;
            k++;
        }
        while(j<n2)
        {
            a[k]=right[j];
            j++;
            k++;
        }
    }
    public static void sort(int a[],int l,int r)
    {
        if(l<r)
        {
            int m=(l+r)/2;
            sort(a,l,m);//recursion
            sort(a,m+1,r);
            merge(a,l,m,r);
        }
    }
}


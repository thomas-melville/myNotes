Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2017-03-07T08:36:46+00:00

###### Interview Questions ######
Created Tuesday 07 March 2017

##### Week 1 #####

#### Question 1 ####


public class QuickUnion implements UnionFind {

	private int[] array;

	public QuickUnion(final int total){
		this.array # new int[total];
		for(int i#0;i<total;i++){
			array[i] # i;
		}
	}

	@Override
	public void union(final int p, final int q) {
		int rootP # root(p);
		int rootQ # root(q);
		array[rootP] # rootQ;
	}

	@Override
	public boolean connected(final int p, final int q) {
		return root(p) ## root(q);
	}
}

#### Question 2 ####

@Override
	public void union(final int p, final int q) {
		int rootP # root(p);
		int rootQ # root(q);
		array[rootP] # rootQ;
		if(q > p){
			largestComponent[p] # q;
		} else {
			largestComponent[q] # p;
		}
	}

	@Override
	public boolean connected(final int p, final int q) {
		return root(p) ## root(q);
	}

	int find(final int i){
		return largestComponent[i];
	}

#### Question 3 ####

public class Hacking {

	private int[] array;
	private int[] successor;

	public Hacking(){
		array # new int[10];
		successor # new int[10];
		for(int i#0;i<10;i++){
			array[i] # i;
			successor[i] # i+1;
		}
		printArray(array);


	}

	public void remove(final int x){
		array[x] # -1;
		successor[x] # successor[x+1];
	}

	public int successor(final int x){
		return successor[x];
	}

	public static void printArray(int[] array) {
		for (int i #0; i< array.length; i++){
			System.out.print("Index: "+i+" "+array[i]);
		}
		System.out.println(" ");
	}
}

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
import java.util.ArrayList;

import org.jsoup.Connection;
import org.jsoup.HttpStatusException;
import org.jsoup.Jsoup;
import org.jsoup.nodes.Document;
import org.jsoup.nodes.Element;
import org.jsoup.select.Elements;
public class TaskCreator 
{
	
	
	
	public static ArrayList<String> expandPatterns(String[] inputArray)
	{
		ArrayList<String> correctedLines = new ArrayList<String>();
		
		for (int x = 0; x<inputArray.length; x++)
		{
			if(inputArray[x].startsWith("!"))//each line if it starts with ! (its a comment so ignore)
			{
				correctedLines.add(inputArray[x]);
			}
			else{
				ArrayList<String> t = URLUtils.expandPatternAndGetUrl(inputArray[x]);
				if(t.size()==0)
					continue;
				else {
					for (String te: t){
						if(URLUtils.isReachable(te, 500))//verify that that page is reachable before adding
							correctedLines.add(te);
						
						
					}
						
			
				}
					
			}
			
		}
		return correctedLines;
		
	}
	
	
	
	
	
	
	
	
	
	/**
	 * Function takes in a url to a page and extracts a url for an image on that page if possible
	 * @param URL: the webpage to scan
	 * @return: the url of smallest possible Img
	 */
	public static String getImgURL(String URL){
		int smallestSize =-1; //this is the smallest dimensioned image val we've encountered yet.
		String bestSizeImgURL = null;
		try {
			if(URL !=null)
			{
				Document theHTMLPage = getPage(URL);
				if(theHTMLPage != null)
				{
					 Elements media = theHTMLPage.select("[src]"); //these are ANY elements with src atr (img or videos)
					 
					 for (Element src : media) 
					 {
				            if (src.tagName().equals("img")) //ensure we are checking an img type element
				            {
				            	String absUrl = src.attr("abs:src"); 
				            	
				            	if(absUrl.contains(URLUtils.getBaseURL(URL)))  //ensure image is in the same domain
				            	{
				            		try //ensure the attributes were set (TODO dimensions arn't set but it IS an image even though we don't know its size. so maybe deal?)
				            		{
					            		int x = Integer.parseInt(src.attr("width"));
					            		int y = Integer.parseInt(src.attr("height"));
					            		
					            		if(smallestSize == -1 || x*y<smallestSize) //found new smallest image
				            			{
					            			smallestSize=x*y; //dimension of smallest image we've found
					            			bestSizeImgURL = src.attr("abs:src");
				            			}
					            		else //new image is larger than current best so ignore it
					            		{}
				            		}
				            		catch (Exception e){
				            			if(bestSizeImgURL == null) { 
				            				bestSizeImgURL = src.attr("abs:src");
				            			}
				            		}
				            		
				            		
		//			                print(String.format("%s: <%s> %sx%s", src.tagName(), src.attr("abs:src"), src.attr("width"), src.attr("height")));
				            	}
				            	
				            	
				            	
				            }
					 }
					return bestSizeImgURL; //no media of correct type
				}
			}
			
			
		} catch (HttpStatusException e) {
			// TODO Auto-generated catch block
			
			e.printStackTrace();
		} 
		
		return null; 
	}
	

	public static void print(String toPrint){
		System.out.println(toPrint);
	}
	
	public static Document getPage(String URL) throws HttpStatusException{
			
			/**
			 * a function that takes in a url and attempts to get the html page to return.
			 */
			Document thePage = null; //the page to return (if request suceeds)
			//TODO when failed connection return error instead of printing error message
			
			Connection.Response response = null;
			try {
				response = Jsoup.connect(URL) //attempt to connect
				        .userAgent("Mozilla")
				        .timeout(10000)
				        .ignoreHttpErrors(true)
				        .execute();
				
			} catch (IOException e1) {
				System.out.println("error establishing connection to website");
			}
	
			int statusCode = response.statusCode(); //observe response status code (200,300,400,500, etc)
	
			if(statusCode == 200) { //good response
				try {
					thePage = response.parse();
				} catch (IOException e) {
					System.out.println("error in parsing the page"); 
					e.printStackTrace();
				}
			}
			if(statusCode >200) { //some error occurred
			
				//System.out.println("error in getting a page (greater than 200 status code returned)" );
				throw new HttpStatusException("issue connecting to: " + URL, statusCode, URL);
				
			}
			
			return thePage;
		}

	
	public static String[] readFile(String filePath){
		ArrayList<String> temp = new ArrayList();
		
		BufferedReader reader;
		try {
			reader = new BufferedReader(new FileReader(filePath));
			String line = reader.readLine();
			while (line != null) {
				temp.add(line);
				line = reader.readLine();
			}
			reader.close();
			
			return temp.toArray(new String[0]);
		} catch (IOException e) {
			e.printStackTrace();
		}
		return null;
	}
	
	public static void main(String[] args) {
		
		String[] ignoreList = {"http://www.filesor.com"};
		
//		//read in our gfwlist 
		String[] fileLines =readFile("shortgfwlist.txt");
//		
//			
		String results[] = expandPatterns(fileLines).toArray(new String[0]); //call expandPatterns to generate urls from apb patterns
//
////		print("got expanded urls " + Integer.toString(results.length));
		write("expandedTest.txt",results);  // write results to a file

		//read in urls that we made
		fileLines =readFile("expandedTest.txt");
		
		
		ArrayList<String> csvLines = new ArrayList<String>();
		for(String x: fileLines)
		{
			try{
			String domain = URLUtils.getBasePageDomain(x);
			print("looping for: " + x); //DEBUGGING
			if(!x.startsWith("!") && !contains(ignoreList, x)) //ignore comments in the url file and the weird entries like "http://www.filesor.com" which don't return pages
			{
				String res = getImgURL(x);
				if(res!=null && (res.endsWith("png") || res.endsWith("jpg") || res.endsWith("jpeg")|| res.endsWith("gif"))) //ensure its a valid image
				csvLines.add(domain+","+res);
			}
			}catch(Exception e){}//silently ignore all errors that may come from getting content (needed for ignoring wierd stuff like "http://www.filesor.com" which don't return status codes and instead return a stream of data in non standard form.
		}
		
		for(String x: csvLines){
			print(x);
		}

		
//		print(URLUtils.getBasePageDomain(url)); //get the domain for a url
//		
//		print(URLUtils.getBaseURL(url)); //get the homepage url
		
	
	}
	
	public static boolean contains(String[] Values, String value)
	{
		for (String s : Values) 
			if (s.equals(value)) 
				return true;
		return false;
	}
	public static void write (String filename, String[] x) {
		  BufferedWriter outputWriter = null;
		  try{
			  outputWriter = new BufferedWriter(new FileWriter(filename));
			  for (int i = 0; i < x.length; i++) {
			    outputWriter.write(x[i]);
			    outputWriter.newLine();
			  }
			  outputWriter.flush();  
			  outputWriter.close();  
		
		  }catch(Exception e){
			  print("couldn't write to file");
		  }
 

		}
	
}

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
		return null;//TODO implement me
		
	}
	
	/**
	 * Function takes in a url to a page and extracts a url for an image on that page if possible
	 * @param URL: the webpage to scan
	 * @return: the url of smallest possible Img
	 */
	public static String getImgURL(String URL){
	
		try {
			Document theHTMLPage = getPage(URL);
			 Elements media = theHTMLPage.select("[src]"); //any elements with src atr (img or videos)
			 
			 for (Element src : media) 
			 {
		            if (src.tagName().equals("img")) //correct type of media
		            {
		            	String absUrl = src.attr("abs:src"); 
		            	if(absUrl.startsWith(URL))  //ensure it fits 
		            	{

			            	//TODO check if size is sufficiently small		            	
			                print(String.format("%s: <%s> %sx%s", src.tagName(), src.attr("abs:src"), src.attr("width"), src.attr("height")));
			            	return src.attr("abs:src");
		            	}
		            	
		            	
		            	
		            }
			 }
			return null; //no media of correct type
			
			
			
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

	
	
	
	public static void main(String[] args) {
		
		String url = "http://www.google.com"; //TODO need to deal with potentially http or https
		getImgURL(url);
		
	
	}
}

import java.net.URL;
import java.net.UnknownHostException;
import java.util.ArrayList;
import java.io.IOException;
import java.net.HttpURLConnection;
import java.net.MalformedURLException;
public class URLUtils {

	
	/**
	 * A helper function to strips a url down to its "base" domain
	 * @param url: the url we want to try to narrow down
	 * @return
	 */
	public static String getBasePageDomain(String url)
	{
		String domain = "";
		

		int b4Domain = url.indexOf("://") +3; //(this is the 3 chars we need to increment by to cut it off completly)
		if (b4Domain > -1) 
		    domain = url.substring(b4Domain); //trim off the protocol
		    
		domain = domain.split("/")[0];
		

		//find & remove port number
		domain = domain.split(":")[0];
		
		return domain;
		
	}
	
	public static String getBaseURL(String urlstr)
	{
		String baseUrl ="";

		try
		{
		  URL url = new URL(urlstr);
		  baseUrl = url.getProtocol() + "://" + url.getHost();
		  return baseUrl;
		}
		catch (MalformedURLException e)
		{
			System.out.println("malformed url: "+ urlstr);
			return baseUrl;
		}
		
	}
	
	
	
	public static ArrayList<String> expandPatternAndGetUrl(String urlPattern)
	{
		ArrayList<String> realURLs = new ArrayList<String>();
		
		//TODO implement me
		/*
		 * if it starts with "|" just trim off "|"
		 * 
		 * if starts with "||"
		 * (technically www. can be any subdomain and any protocol)
			 * prepend http:// 		(and add to real URLs)
			 * prepend: http://www. (and add to real URLs)
		
		 * if it starts with: "@@|" or "@@||"
		 	* indicates that any subdomain except the listed domain is acceptable
		 	* example: @@|http://8.example.org means http://*.example.org (with any value for star) is blocked
		 	* since we don't want to deal with this subdomain expansion. just drop the URL
		 	* 
		 *
		 */
		
		if(urlPattern.startsWith("||")){
			String t = "http://" + urlPattern.substring(2);
			realURLs.add(t);
			t = "http://www." + urlPattern.substring(2);
			realURLs.add(t);
			//TODO exand with other subdomains	
			t = "https://" + urlPattern.substring(2);
			realURLs.add(t);
			t = "https://www." + urlPattern.substring(2);
			realURLs.add(t);
			
			
		}
		else if (urlPattern.startsWith("|")){
			realURLs.add( urlPattern.substring(1));
		}
		else if (urlPattern.startsWith("@@"))
		{
			//TODO implement this once we have broader range of potential subdomains it could match
		}
		else if (urlPattern.startsWith(".")) //no idea what this is (their documentation doesn't explain but I think it ist just matching the domain, for all possible subdomains)
		{
			String t = "http://www" + urlPattern;
			realURLs.add(t);
		}
		return realURLs;
		
	}

	public static boolean isReachable(String targetUrl, int timeout) 
	{
	    HttpURLConnection httpUrlConnection;
		try {
			httpUrlConnection = (HttpURLConnection) new URL(targetUrl).openConnection();
			   httpUrlConnection.setRequestMethod("HEAD"); //USING HEAD request instead of getting entire page to save on bandwidth usage
			   httpUrlConnection.setConnectTimeout(timeout);
			   httpUrlConnection.setReadTimeout(timeout);
	
			   
			   try
			    {
			        int responseCode = httpUrlConnection.getResponseCode();
			        
			        return responseCode == HttpURLConnection.HTTP_OK; //we could connect
			    } catch (UnknownHostException noInternetConnection)
			    {
			    	
			        return false;
			    }
			
		} catch (MalformedURLException e) {
			//we couldn't connect
			System.out.println(" had error with: " + targetUrl);
		} catch (IOException e) {
			// our device is having network issues 
			System.out.println("had error with: " + targetUrl);
		}
		return false;
	 

	  
	}
	

}

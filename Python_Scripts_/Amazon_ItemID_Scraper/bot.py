import os
import selenium
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
import time
import discord
from discord.ext import commands
from discord.ext.commands import has_permissions
from discord.utils import get
from dotenv import load_dotenv

bot = commands.Bot(command_prefix='!')

load_dotenv()
TOKEN = os.getenv('DISCORD_TOKEN')

@bot.event
async def on_ready():
    print(f'{bot.user} is connected to the following guilds:')

    for guild in bot.guilds:
        print(f'{guild.name} (id: {guild.id})')

# GetID command
@bot.command(name='getid', help='Gets Amazon Offer ID')
async def getid(ctx, link):
    # Sets up Selenium webdriver to open link without browser GUI
    chrome_options = webdriver.ChromeOptions()
    chrome_options.add_argument("--headless")
    
    driver = webdriver.Chrome(chrome_options=chrome_options)
    driver.get(link)
    
    # Attempts to extract ASIN from link
    asin = link.split('?')
    asin = asin[0].split('/')
    print(asin)

    asinno = ""
    # Error handling, in case there is a / at the end of the link
    if asin[-1] == "":
        print(asin[-2])
        asinno = asin[-2]
    else:
        print(asin[-1])
        asinno = asin[-1]
    # Stores item name from browser tab title
    item = driver.title
    print(item)
    
    # Finds and clicks add to cart button
    addcart = driver.find_element_by_id("add-to-cart-button")
    addcart.click()
    
    # Wait 5 seconds, ensures item added to cart
    time.sleep(5)
    
    # Switches to cart tab
    driver.get("https://www.amazon.com/gp/cart/view.html?ref_=nav_cart")
    
    # Writes cart HTML to buffer.txt 
    # yes, I know this isn't the most efficient way to do this but it seemed the simplest at the time
    html = driver.page_source
    with open('buffer.txt', 'w') as f:
        f.write(html)
    
    # Iterates through code lines looking for the line of HTML which includes the ASIN and subsequently the OfferID
    # Starts at 3000 since the HTML line we want is always below line 3000
    line_number = 3000
    list_of_results = []

    with open('buffer.txt', 'r') as read_obj:
        line_number = 3000
        list_of_results = []
        for line in read_obj:
            line_number += 1
            if asinno in line:
                list_of_results.append((line_number, line.rstrip()))
    
    # Selects the desired HTML line and strips the rest of the line to leave only the OfferID
    htmlline = list_of_results[0]
    htmlline = htmlline[1].split('data-encoded-offering="')
    htmlline = htmlline[1]
    htmlline = htmlline.split('" ')
    OfferID = htmlline[0]
    await ctx.send("Offer ID for "+ item +"\n" + OfferID)
    
    # Closes the browser after returning the OfferID
    driver.quit()

bot.run(TOKEN)

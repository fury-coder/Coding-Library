# amazon-discord-bot
A Discord.py bot that, when provided the link to an Amazon listing, extracts the item's OfferID (also known as data-encoded-offering) and delivers it in a message in the Discord channel.

# installation
Download the files, create a Discord bot application at https://discord.com/developers/applications, then paste the bot token in the .env file

# usage
`!getid <link>` to get the OfferID for the linked item.

**NOTE:** As Amazon links can be rather inconsistent, this works best with links formatted such that the ASIN is the last item in the link. For example:

https://www.amazon.com/Never-Gonna-Give-You-Up/dp/B07X66DCLM

as opposed to

https://www.amazon.com/Never-Gonna-Give-You-Up/dp/B07X66DCLM/ref=sr_1_1?dchild=1&keywords=never+gonna+give+you+up&qid=1620572061&sr=8-1

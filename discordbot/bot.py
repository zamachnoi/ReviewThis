from dotenv import load_dotenv
import discord
import os
from discord import app_commands

load_dotenv()

token = os.getenv("TOKEN")
intents = discord.Intents.all()
client = discord.Client(intents=intents)
tree = app_commands.CommandTree(client)

@client.event
async def on_ready():
    await tree.sync()

@tree.command(name="setqueue", description="Set the queue for the server")
async def setqueue(interaction: discord.Interaction):
    await interaction.response.send_message("Setting queue")



client.run(token)
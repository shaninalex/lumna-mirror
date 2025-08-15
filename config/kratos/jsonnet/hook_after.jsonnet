function(ctx) {
  userId: ctx.identity.id,
  traits: {
    email: ctx.identity.traits.email,
  },
}
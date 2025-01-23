import { ethers } from "hardhat";

async function transferWithMemo() {
  const classId = "0x7f2a8B018075A412bE100EFF15b0F3E4c6DE96B4";
  const tokenId = 0;
  const from = "0x98Aa0E7E441B616b5291E67b38A09c32425DE6eA";
  const to = "0xD9777d58b4725738900181D077B59C3B4BBBCbaD";

  const signer = await ethers.provider.getSigner();

  const Class = await ethers.getContractFactory("Class", {
    signer,
  });

  const class_ = Class.attach(classId);

  const tx = await class_.transferWithMemo(
    from,
    to,
    tokenId,
    `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin quis semper risus. Pellentesque in ante nulla. Donec in nisi eu urna euismod condimentum. Integer in dui odio. Phasellus vel enim ac elit auctor congue. Morbi vitae consequat dolor, ultrices tristique nulla. Mauris consequat arcu nec lorem vestibulum, quis auctor mi iaculis. Curabitur viverra vulputate tristique. Sed dignissim, eros eget sollicitudin tincidunt, metus eros interdum eros, id vulputate lectus tortor id ipsum.

Vivamus egestas mi felis, et semper eros malesuada vitae. Cras quis eros rhoncus, eleifend elit sit amet, sagittis neque. Sed non neque sapien. Nullam ultricies commodo felis, nec suscipit mauris. Nam nec diam molestie, vestibulum eros ac, egestas velit. Aliquam lobortis justo eget blandit imperdiet. Aliquam facilisis sapien quis iaculis auctor. Ut et malesuada arcu, at posuere velit. In eget sem vulputate, dignissim risus ut, eleifend nunc. Praesent sit amet justo id tortor convallis accumsan quis sit amet lorem. Sed purus risus, ullamcorper a elit nec, bibendum posuere orci.

Suspendisse at nulla suscipit, semper leo quis, fringilla ante. Fusce et tortor scelerisque enim efficitur volutpat non id lectus. Morbi congue justo vitae justo suscipit volutpat. Suspendisse auctor ex augue. Quisque eleifend massa sit amet fermentum tristique. Fusce quis neque accumsan, interdum tortor sit amet, tempus lorem. Praesent ut eleifend dui, at auctor nisi. Nam hendrerit dictum hendrerit. Ut tristique, ipsum sed lobortis dignissim, arcu erat mollis dui, ut sodales justo quam in leo. Duis a libero eu nulla placerat laoreet. Proin euismod id risus in egestas. Donec id felis ut sem fermentum ullamcorper. Etiam tempor, sapien eget vestibulum lobortis, metus tellus congue nisl, a fermentum ipsum nisi at arcu. Sed vel nisi non enim auctor dictum.

Donec quis sagittis nibh, ut iaculis mauris. Sed ultricies sodales lectus, sed volutpat tortor aliquam a. Aenean luctus ligula at molestie rhoncus. Vivamus iaculis dictum dui, id aliquet erat maximus a. Sed turpis elit, condimentum ac risus vel, bibendum tristique ante. Cras eu pretium ante, non tincidunt turpis. Vestibulum a nibh sit amet ipsum pharetra consequat ac vel elit. Donec efficitur nisi nec aliquet aliquet.

Duis nec augue neque. Etiam viverra at magna ac ornare. Curabitur et eleifend ante. Vestibulum interdum consectetur ante nec pharetra. Maecenas sagittis libero libero, a consectetur nunc interdum quis. Etiam feugiat euismod enim, quis efficitur erat gravida pulvinar. Fusce mollis bibendum quam, quis luctus arcu facilisis ac. Nullam ultrices maximus mi, quis condimentum nisl posuere nec. Suspendisse pretium nisi erat, sed ultrices sem facilisis in. Sed vel justo dui. Morbi scelerisque ex et ultrices sagittis.`,
  );
  console.log(await tx.wait());
}

transferWithMemo().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
